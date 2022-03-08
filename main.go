package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	connectDB()
	router.POST("/", userReg)
	router.POST("/login", login)
	// router.GET("/getDatabyId/:id", getDatabyID)
	router.Run("localhost: 5000")
}
