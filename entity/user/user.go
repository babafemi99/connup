package user

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Role int

const (
	Admin Role = iota
	Member
)

type User struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	Role        Role      `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Deactivated bool      `json:"deactivated"`
	IsVerified  bool      `json:"is_verified"`
}
type AccessUser struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Role        Role   `json:"role"`
	Deactivated bool   `json:"deactivated"`
	IsVerified  bool   `json:"is_verified"`
}

func (u *User) ComparePasswords(s string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(s))
}
