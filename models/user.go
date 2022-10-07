package models

// BUILD USER MODEL
type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `gorm:"unique"`
	Password []byte `json:"-"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
