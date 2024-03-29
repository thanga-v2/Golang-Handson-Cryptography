package main

import "fmt"

func main() {
	// declare variables
	var str string
	var n, m int
	var mn float32

	// assign values
	str = "Hello thanga"
	n = 10
	m = 50
	mn = 2.45

	fmt.Println("value of str= ", str)
	fmt.Println("value of n= ", n)
	fmt.Println("value of m= ", m)
	fmt.Println("value of mn= ", mn)

	// declare and assign values to variables
	var city string = "Mumbai"
	var x int = 100

	fmt.Println("value of city= ", city)
	fmt.Println("value of x= ", x)

	// declare variable with defining its type
	country := "IND"
	val := 15
	fmt.Println("value of country= ", country)
	fmt.Println("value of val= ", val)

	// define multiple variables
	var (
		name  string
		email string
		age   int
	)
	name = "thanga"
	email = "thangarajm@icloud.com"
	age = 25

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)
}
