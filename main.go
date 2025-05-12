// package main

// import (
// 	"fmt"
// 	// "time"
// )

// //	type PERSON struct {
// //	 Name string
// //	 Age  int
// //	}
// func temp() {
// 	// 	// var name string = "Ishpreet"
// 	// 	// fmt.Printf("This is my name %s\n ", name)
// 	// 	// age := 19 // shorthand operator
// 	// 	// fmt.Printf("This is my age %d\n", age)
// 	// 	// var city string
// 	// 	// city = "Delhi"
// 	// 	// fmt.Printf("This is my city %s\n", city)

// 	// 	// var (
// 	// 	// 	isEmployed bool   = true
// 	// 	// 	salary     int    = 11000
// 	// 	// 	position   string = "developer"
// 	// 	// )
// 	// 	// fmt.Printf("isemployes %t at position : %s with salary : %d\n", isEmployed, position, salary)
// 	// 	// // Zero values
// 	// 	// var defaultint int

// 	// 	// fmt.Printf("int %d\n", defaultint)
// 	// 	// const pi = 3.14
// 	// 	// fmt.Printf("Pi : %d\n", pi)

// 	// 	// const (
// 	// 	// 	Jan int = iota + 1 // enum ka jugad
// 	// 	// 	Feb
// 	// 	// 	Mar
// 	// 	// 	Apr
// 	// 	// )
// 	// 	// fmt.Printf("%d%d%d%d", Jan, Feb, Mar, Apr)
// 	// 	// result := add(3, 4)
// 	// 	// fmt.Printf("\nThis is result %d", result)
// 	// 	// sum, _ := sumproduct(3, 4)
// 	// 	// _, product := sumproduct(3, 4)
// 	// 	// fmt.Printf("\nThis is sum %d", sum)
// 	// 	// fmt.Printf("\nThis is product %d", product)
// 	// 	// 	age:= 19
// 	// 	// 	if age >= 18 {
// 	// 	// 		fmt.Println("Adult")
// 	// 	// 	}
// 	// 	// }
// 	// 	// func add(a int, b int) int {
// 	// 	// 	return a + b
// 	// 	// }
// 	// 	// func sumproduct(a, b int) (int, int) {
// 	// 	// 	return a + b, a * b
// 	// 	// number := [4] int { 1,2,3,4}
// 	// 	// fmt.Printf("%v/n",number)
// 	// 	// fruits := []string{"apple","banana"}
// 	// 	// fmt.Printf("%v",fruits)
// 	// 	// fruits= append(fruits, "strawberry")
// 	// 	// fmt.Printf("%v",fruits)
// 	// 	// for _,value := range number{
// 	// 	// 	fmt.Printf("%d",value)
// 	// 	// }
// 	// 	// capCities := map[string]string{  // MAP
// 	// 	// 	"INDIA": "DELHI",
// 	// 	// 	"USA":   "WASHINGTON DC",
// 	// 	// }
// 	// 	// fmt.Printf(capCities["INDIA"])
// 	// 	// person := PERSON{Name :"Ishpreet",Age : 19 }
// 	// 	// fmt.Printf("Person : %+v",person)
// }
// func main() {
// 	// fmt.Println("Enter your name:")
// 	// reader := bufio.NewReader(os.Stdin)
// 	// input, err := reader.ReadString('\n')
// 	// fmt.Printf("hello %s", input)
// 	// fmt.Printf("rate our burger from 1 to 5")
// 	// reader := bufio.NewReader(os.Stdin)
// 	// input, _ := reader.ReadString('\n')
// 	// fmt.Printf("You rated %s", input)
// 	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
// 	// if err != nil {
// 	// 	fmt.Printf("print error", err)
// 	// } else {
// 	// 	fmt.Println(numRating + 1)
// 	// }
// 	// presentTime := time.Now()
// 	// fmt.Println(presentTime)
// 	// fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))
// 	defer fmt.Printf("this is defer statement")
// 	defer fmt.Printf("one more defer statement") //LIFO
// 	fmt.Printf("this will be printed first \n ")

// }

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/me/:id", func(c *gin.Context) {
		var id = c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"my_id":   id,
		})
	})
	r.POST("/post", func(c *gin.Context) {
		type postreq struct {
			Email    string `json:"email"'`
			Password string `json:"password"`
		}
		var request postreq
		if err:= c.BindJSON(&request) ; err!=nil{
			c.JSON(http.StatusBadRequest , gin.H{
				"error":err.Error(),

			})
			return
		}
		c.BindJSON(&request)
		c.JSON(http.StatusOK, gin.H{
			"email":    request.Email,
			"password": request.Password,
		})

	})
	r.Run(":8000")
}
