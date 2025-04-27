package main

import (
	"fmt"
)

func main() {
	var name string = "Ishpreet"
	fmt.Printf("This is my name %s\n ",name)
	age:= 19   // shorthand operator
	fmt.Printf("This is my age %d\n",age)
	var city string
	city = "Delhi"
	fmt.Printf("This is my city %s\n",city)

	var (
		isEmployed bool =true
		salary int = 11000
		position string = "developer"
	)
	fmt.Printf("isemployes %t at position : %s with salary : %d\n",isEmployed,position,salary)
	// Zero values
	var defaultint int

	fmt.Printf("int %d",defaultint)
}
