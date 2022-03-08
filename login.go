package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func login(c *gin.Context) {
	var userCollection = GetCollection(DB, "myUsers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	credentials := new(User)
	defer cancel()

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}

	userLoginDetails := User{
		Username: credentials.Username,
		Password: credentials.Password,
	}

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var userInfo []bson.M
	if err = cursor.All(ctx, &userInfo); err != nil {
		log.Fatal(err)
	}

	var UserConfirmed bool = false

	// Authenticating user credentials and checking password hash:

	for i := 0; i < len(userInfo); i++ {
		if userInfo[i]["Username"].(string) == userLoginDetails.Username &&
			CheckPasswordHash(userLoginDetails.Password, userInfo[i]["Password"].(string)) {
			UserConfirmed = true
			break

		}

	}

	if UserConfirmed {
		c.JSON(http.StatusCreated, gin.H{"message": "User authenticated!"})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid credentials"})
		return
	}

}
