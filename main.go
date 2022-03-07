package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Arraytest struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	First_name string             `bson:"First_name,omitempty bson:"First_name,omitempty"`
	Last_name  string             `bson:"Last_name,omitempty bson:"Last_name,omitempty"`
	Username   string             `bson:"Username,omitempty bson:"Username,omitempty"`
	Email      string             `bson:"Email,omitempty bson:"Email,omitempty"`
	Password   string             `bson:"Password,omitempty bson:"Password,omitempty"`
}

// var emp1 = Arraytest{Name: "Idowu", Salary: 12345.09}

// var client *mongo.Client

var DB = connectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("myGoappDB").Collection("myUsers")
	return collection
}

func postData(c *gin.Context) {
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
		Password:   HashPassword(test.Password),
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
	if len(newTest.Password) < 8 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Password must be longer than 8 character"})
		return
	}

	if res {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "This email exists"})
		return
	} else if userRes {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "This username exists"})
		return
	}

	result, err := userCollection.InsertOne(ctx, newTest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success", "Data": map[string]interface{}{"data": result}})

	// c.IndentedJSON(http.StatusOK, gin.H{"message": result})

}

// func postData(c *gin.Context) {
// 	var newData Arraytest

// 	if err := c.BindJSON(&newData); err != nil {
// 		return
// 	}

// 	emp1 = append(emp1, newData)
// 	c.IndentedJSON(http.StatusCreated, newData)

// }

// func getDatabyID(c *gin.Context) {
// 	name := c.Param("name")

// 	for i := 0; i < len(emp1); i++ {
// 		if emp1[i].Name == name {
// 			c.IndentedJSON(http.StatusOK, emp1[i])
// 			return

// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
// }

func main() {

	router := gin.Default()
	connectDB()
	router.POST("/", postData)
	// router.POST("/postData", postData)
	// router.GET("/getDatabyId/:id", getDatabyID)
	router.Run("localhost: 5000")
}

// func main() {
// 	fmt.Println("Hello world")
// 	const Name = "Idowu"
// 	var file = "Paul"
// 	// myArray = [1, 2, 4]
// 	fmt.Println(Name + " " + file)
// 	fmt.Println("What is your name:")

// 	var name string
// 	fmt.Scanln(&name)
// 	fmt.Println("Welcome,", name)
// 	var A = false
// 	println(A)

// type Arraytest struct {
// 	ID     int
// 	Name   string
// 	Salary float32
// }

// var emp1 = []Arraytest{
// 	{ID: 1, Name: "Idowu", Salary: 12345.09},
// 	{ID: 2, Name: "Paul", Salary: 300025.09},
// }

// var p = []string{}

// for i := range emp1 {
// 	p := append(p, emp1[i].Name)
// 	fmt.Print(p)

// 	// fmt.Println(emp1[i].Name)
// }

// for i := 0; i < len(emp1); i++ {
// 	p := append(p, emp1[i].Name)
// 	fmt.Println(p)
// 	fmt.Println(emp1[i].Name)
// }

// fmt.Println(emp1[0].ID)

// // fmt.Println(&Arraytest)

// var myArray = []int{1, 3, 5, 6}
// // myArray2 := []string{"Idowu", "Paul"}

// fmt.Println(myArray)

// aMap := map[string]string{
// 	"Name":   "Idowu",
// 	"Status": "Single",
// 	"Lines":  "Jog",
// }
// fmt.Println(aMap["Name"])

// // for i := 0; i < len(aMap); i++ {
// // 	fmt.Println(aMap[i])

// // }

// for i := range aMap {
// 	fmt.Println(aMap[i])
// }

// for i := 0; i < len(aMap); i++ {
// 	fmt.Println(i)
// }

// for i := range Arraytest{

// }

// fmt.Println(myArray2)
// fmt.Println(myArray)

// for i := 0; i < myArray.count; i++ {

// }
// fmt.Println(myArray)
// }
