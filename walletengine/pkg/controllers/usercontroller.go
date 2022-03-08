package controllers

import (
	"RPay/walletengine/pkg/models"

	"github.com/gin-gonic/gin"
)

var NewUser models.User

func GetUsers(c *gin.Context) {
	users := models.GetAllUsers()
	c.IndentedJSON(200, users)
}

func AddUser(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	user = models.AddNewUser(user)
	c.IndentedJSON(200, user)
}
