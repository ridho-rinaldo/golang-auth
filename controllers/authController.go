package controllers

import (
	"golang-auth/database"
	"golang-auth/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
		Gender:   data["gender"],
		Phone:    data["phone"],
		Address:  data["address"],
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"status":  201,
		"message": "Registrasi Sukses",
		"data":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Akun tidak ditemukan",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Password invalid",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Proses Login Gagal",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Sukses",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		var data []string
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": "Tidak diizinkan",
			"data":    data,
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Sukses",
		"data":    &user,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Logout sukses",
	})
}

func Update(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		var null []string
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": "Tidak diizinkan",
			"data":    null,
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	var password []byte
	if data["password"] == "" {

		database.DB.
			Where("id = ?", claims.Issuer).
			Model(&user).
			Updates(models.User{
				Username: data["username"],
				Email:    data["email"],
				Gender:   data["gender"],
				Phone:    data["phone"],
				Address:  data["address"],
			})
	} else {
		decrypt, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
		password = decrypt

		database.DB.
			Where("id = ?", claims.Issuer).
			Model(&user).
			Updates(models.User{
				Username: data["username"],
				Email:    data["email"],
				Password: password,
				Gender:   data["gender"],
				Phone:    data["phone"],
				Address:  data["address"],
			})
	}

	return c.JSON(fiber.Map{
		"status":  201,
		"message": "Update Sukses",
		"data":    user,
	})
}

func GetAll(c *fiber.Ctx) error {

	allUser := []models.User{}

	database.DB.Find(&allUser)

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Sukses Ambil Data",
		"data":    &allUser,
	})
}
