package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	getcollection "user/Desktop/Desktop/go_code/Collection"
	database "user/Desktop/Desktop/go_code/databases"
	passwordhash "user/Desktop/Desktop/go_code/hashpassword"
	model "user/Desktop/Desktop/go_code/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *gin.Context) {
	var DB = database.ConnectDB()
	var userCollection = getcollection.GetCollection(DB, "myUsers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	credentials := new(model.User)
	defer cancel()

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}

	userLoginDetails := model.User{
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
			passwordhash.CheckPasswordHash(userLoginDetails.Password, userInfo[i]["Password"].(string)) {
			UserConfirmed = true
			break
		}
	}

	if UserConfirmed {
		c.JSON(http.StatusCreated, gin.H{"message": "User authenticated!"})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid credentials. Maybe you need to sign up"})
		return
	}

}
