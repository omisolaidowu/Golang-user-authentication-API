package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	getcollection "go_code/Collection"
	database "go_code/databases"
	passwordhash "go_code/hashpassword"
	model "go_code/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserReg(c *gin.Context) {
	var DB = database.ConnectDB()
	var userCollection = getcollection.GetCollection(DB, "myUsers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	user := new(model.User)
	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	userPayload := model.User{
		ID:         primitive.NewObjectID(),
		First_name: user.First_name,
		Last_name:  user.Last_name,
		Username:   user.Username,
		Email:      user.Email,
		Password:   user.Password,
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

	var emailExists bool = false
	var usernameExists bool = false

	for i := 0; i < len(userInfo); i++ {
		if userInfo[i]["Email"].(string) == userPayload.Email {
			emailExists = true
			break
		} else if userInfo[i]["Username"].(string) == userPayload.Username {
			usernameExists = true
			break
		}
	}

	if emailExists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "This email exists"})
		return
	} else if usernameExists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "This username exists"})
		return
	} else if len(userPayload.Password) < 8 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Password must be longer than 8 character"})
		return
	}
	// Pass the hashed password into the payload here:
	// Reason: hashing the password directly within the payload doesn't throw
	// correct error since hashed password is about 60 to 100 character long
	userPayload.Password = passwordhash.HashPassword(userPayload.Password)

	result, err := userCollection.InsertOne(ctx, userPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success", "Data": map[string]interface{}{"data": result}})
}
