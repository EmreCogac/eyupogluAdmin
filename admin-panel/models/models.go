package models

import (
	"admin-panel/admin-panel/database"
	"time"

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
	Title    string `gorm:"type:varchar(30);not null" binding:"required"`
	Location string `gorm:"type:varchar(30);not null" binding:"required"`
	Type     string `gorm:"type:varchar(30);not null" binding:"required"`
	Info     string `gorm:"type:varchar(40);not null" binding:"required"`
	State    string `gorm:"type:varchar(30);not null" binding:"required"`
	Img      []byte `json:"picture"`
}

func (now *Ilanlar) UpdatePost(updateted *Ilanlar, id int) error {

	db := database.GlobalDB
	now.UpdatedAt = time.Now()
	now.Title = updateted.Title
	now.Location = updateted.Location
	now.Type = updateted.Type
	now.Info = updateted.Info
	now.State = updateted.State
	now.Img = updateted.Img
	err := db.Where("id= ?", id).Updates(&now)
	if err != nil {
		return err.Error
	}
	return nil

}

func (deleted *Ilanlar) DeletePost(id int) error {
	db := database.GlobalDB

	err := db.Unscoped().Where("id =? ", id).Delete(&deleted)
	if err != nil {
		return err.Error
	}
	return nil
}

func (post *Ilanlar) ModelCreatePost() error {
	result := database.GlobalDB.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil

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
