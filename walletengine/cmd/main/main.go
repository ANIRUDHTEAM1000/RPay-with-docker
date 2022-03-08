package main

import (
	"RPay/walletengine/pkg/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
