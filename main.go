package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	connectDB()
	router.POST("/", userReg)
	// router.POST("/postData", postData)
	// router.GET("/getDatabyId/:id", getDatabyID)
	router.Run("localhost: 5000")
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
