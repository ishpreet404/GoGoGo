package main

import (
	"fmt"
)

func main() {
	var name string = "Ishpreet"
	fmt.Printf("This is my name %s\n ", name)
	age := 19 // shorthand operator
	fmt.Printf("This is my age %d\n", age)
	var city string
	city = "Delhi"
	fmt.Printf("This is my city %s\n", city)

	var (
		isEmployed bool   = true
		salary     int    = 11000
		position   string = "developer"
	)
	fmt.Printf("isemployes %t at position : %s with salary : %d\n", isEmployed, position, salary)
	// Zero values
	var defaultint int

	fmt.Printf("int %d\n", defaultint)
	const pi = 3.14
	fmt.Printf("Pi : %d\n", pi)

	const (
		Jan int = iota + 1 // enum ka jugad
		Feb
		Mar
		Apr
	)
	fmt.Printf("%d%d%d%d", Jan, Feb, Mar, Apr)
	result := add(3, 4)
	fmt.Printf("\nThis is result %d", result)
	sum, _ := sumproduct(3, 4)
	_, product := sumproduct(3, 4)
	fmt.Printf("\nThis is sum %d", sum)
	fmt.Printf("\nThis is product %d", product)
}
func add(a int, b int) int {
	return a + b
}
func sumproduct(a, b int) (int, int) {
	return a + b, a * b
}
