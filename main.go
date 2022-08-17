package main

import (
	routes "go_code/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/", routes.UserReg)
	router.POST("/login", routes.Login)
	// router.GET("/getDatabyId/:id", getDatabyID)
	router.Run("localhost: 8000")
}
