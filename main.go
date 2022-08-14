package main

import (
	database "user/Desktop/Desktop/go_code/databases"
	routes "user/Desktop/Desktop/go_code/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	database.ConnectDB()
	router.POST("/", routes.UserReg)
	router.POST("/login", routes.Login)
	// router.GET("/getDatabyId/:id", getDatabyID)
	router.Run("localhost: 5000")
}
