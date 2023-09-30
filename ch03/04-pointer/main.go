package main

import (
	"fmt"
	"reflect"
)

func createPointer() *float64 {
	myFloat := 98.5

	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func double(num *int) {
	*num *= 2
}

func main() {
	amount := 6

	fmt.Println(amount)
	fmt.Println(&amount)

	fmt.Println()

	var myInt int = 4

	fmt.Println(&myInt)
	fmt.Println(reflect.TypeOf(&myInt))

	fmt.Println()

	// var myIntPointer *int
	// myIntPointer = &myInt

	myIntPointer := &myInt

	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	fmt.Println()

	*myIntPointer = 8 

	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)

	fmt.Println()

	myFloatPointer := createPointer()

	fmt.Println(*myFloatPointer)

	fmt.Println()

	myBool := true

	printPointer(&myBool)

	fmt.Println()

	double(&amount)

	fmt.Println(amount)
}