package models

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null;type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	log.Printf("DEBUG: Password hashed to: %s", u.Password)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	log.Printf("DEBUG: Comparing password with hash: %s", u.Password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Printf("DEBUG: Password comparison failed: %v", err)
		return false
	}
	log.Printf("DEBUG: Password comparison successful")
	return true
}
