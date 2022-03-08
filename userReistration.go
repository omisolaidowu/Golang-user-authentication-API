package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DB = connectDB()

func userReg(c *gin.Context) {
	var userCollection = GetCollection(DB, "myUsers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	test := new(Arraytest)
	defer cancel()

	if err := c.BindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	newTest := Arraytest{
		ID:         primitive.NewObjectID(),
		First_name: test.First_name,
		Last_name:  test.Last_name,
		Username:   test.Username,
		Email:      test.Email,
		Password:   test.Password,
	}
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var userInfo []bson.M
	if err = cursor.All(ctx, &userInfo); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(userInfo)

	var emails []string
	var usernames []string

	for i := 0; i < len(userInfo); i++ {
		emails = append(emails, userInfo[i]["Email"].(string))
		usernames = append(emails, userInfo[i]["Username"].(string))
	}

	var res bool = false
	var userRes bool = false
	for _, x := range emails {
		if x == newTest.Email {
			res = true
			break
		}
	}

	for _, x := range usernames {
		if x == newTest.Username {
			userRes = true
			break
		}
	}

	if res {
		fmt.Println(len(newTest.Password))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "This email exists"})
		return
	} else if userRes {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "This username exists"})
		return
	} else if len(newTest.Password) < 8 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Password must be longer than 8 character"})
		return
	}
	// Pass the hashed password into the payload here:
	// Reason: hashing the password directly within the payload doesn't throw
	// correct error since hashed password is about 60 to 100 character long
	newTest.Password = HashPassword(newTest.Password)

	result, err := userCollection.InsertOne(ctx, newTest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success", "Data": map[string]interface{}{"data": result}})
}
