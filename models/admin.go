package models

import "golang.org/x/crypto/bcrypt"

type Admin struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	Phone    string `json:"phone"`
}

func (admin *Admin) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	admin.Password = hashedPassword
}

func (admin *Admin) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(admin.Password, []byte(password))
}
