package models

import (
	"admin-panel/admin-panel/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type Ilanlar struct {
	gorm.Model
	ID       uint8  `gorm:"primaryKey"`
	Location string `gorm:"type:varchar(30);not null"`
	Type     string `gorm:"type:varchar(30);not null"`
	Info     string `gorm:"type:varchar(30);not null"`
	State    string `gorm:"type:varchar(30);not null"`
	Img      []byte `json:"picture"`
}

func (user *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
