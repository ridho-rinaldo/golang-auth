package controllers

import (
	"golang-auth/database"
	"golang-auth/models"
	"strconv"
	"time"

	// JSON WEB TOKEN (JWT) AUTH LIBRARY
	"github.com/dgrijalva/jwt-go"
	// FRAMEWORK WEB APP, SAMA SEPERTI EXPRESSJS DI NODE JS
	"github.com/gofiber/fiber/v2"
	// ENCRYPT DECRYPT LIBRARY -> BIASANYA UNTUK PASSWORD
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// CHECK THE REQUEST
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	// ENCRYPT PASSWORD
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
		Gender:   data["gender"],
		Phone:    data["phone"],
		Address:  data["address"],
	}

	// SIMPAN DATA KE DB
	database.DB.Create(&user)

	// RETURN JSON RESPONSE
	return c.JSON(fiber.Map{
		"status":  201,
		"message": "Registrasi Sukses",
		"data":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	// CHECK THE REQUEST
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	var user models.User

	// CHECK USER IN DB
	database.DB.Where("email = ?", data["email"]).First(&user)

	// RESPONSE WHEN ACCOUNT NOT FOUND
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Akun tidak ditemukan",
		})
	}

	// RESPONSE WHEN INVALID PASSWORD
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Password invalid",
		})
	}

	// PREPARE CREATE NEW TOKEN
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	})

	// GET COMPLETE SIGNED TOKEN
	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Proses Login Gagal",
		})
	}

	// PREPARE SAVE TOKEN TO COOKIES
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	// SAVE AUTHENTICATION TOKEN TO COOKIES
	c.Cookie(&cookie)

	// RETURN JSON RESPONSE
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Sukses",
	})
}

func User(c *fiber.Ctx) error {

	// GET COOKIES WITH PARAM
	cookie := c.Cookies("jwt")

	// PARSE, VALIDATE, AND RETURN A TOKEN.
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// CONDITION WHEN TOKEN UNAUTHORIZED
	if err != nil {
		var data []string
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": "Tidak diizinkan",
			"data":    data,
		})
	}

	// RECHECK TOKEN AND RETURN TO VAR CLAIMS
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	// QUERY SELECT DB
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	// RETURN JSON RESPONSE
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Sukses",
		"data":    &user,
	})
}

func Logout(c *fiber.Ctx) error {

	// PREPARE TO RESET TOKEN
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	// SAVE TO COOKIES
	c.Cookie(&cookie)

	// RETURN JSON RESPONSE
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Logout sukses",
	})
}

func Update(c *fiber.Ctx) error {
	var data map[string]string

	// CHECK THE REQUEST
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	// GET COOKIES WITH PARAM
	cookie := c.Cookies("jwt")

	// PARSE, VALIDATE, AND RETURN A TOKEN.
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// CONDITION WHEN TOKEN UNAUTHORIZED
	if err != nil {
		var null []string
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": "Tidak diizinkan",
			"data":    null,
		})
	}

	// RECHECK TOKEN AND RETURN TO VAR CLAIMS
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	var password []byte

	// SEPARATE CONDITION IS UPDATING PASSWORD OR NOT
	if data["password"] == "" { // PASSWORD NOT UPDATED

		// UPDATE USER TO DATABASE
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
	} else { // PASSWORD UPDATED

		// ENCRYPT NEW PASSWORD
		decrypt, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
		password = decrypt

		// UPDATE USER TO DATABASE
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

	// RETURN JSON RESPONSE
	return c.JSON(fiber.Map{
		"status":  201,
		"message": "Update Sukses",
		"data":    user,
	})
}

func GetAll(c *fiber.Ctx) error {

	allUser := []models.User{}

	// GET ALL USERS
	database.DB.Find(&allUser)

	// RETURN JSON RESPONSE
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Sukses Ambil Data",
		"data":    &allUser,
	})
}
