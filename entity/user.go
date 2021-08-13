package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string
	Password string
}

func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}