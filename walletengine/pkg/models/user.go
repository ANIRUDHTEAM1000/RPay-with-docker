package models

import (
	"fmt"

	"RPay/walletengine/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	UserName    string `json:"UserName" binding:"required"`
	UserId      int    `json:"UserId"  binding:"required" `
	DateOfBirth string `json:"DateOfBirth"  binding:"required" `
	PhoneNumber int    `json:"PhoneNumber"  binding:"required" `
	EmailId     string `json:"EmailId"  binding:"required" `
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func AddNewUser(user User) User {
	fmt.Println("In add new user", user)
	db.Save(&user)
	return user
}

// func GetUserById(id int) User {
// 	var user User
// 	db.First(&user, id)
// 	return user
// }
