package main

import (
	"fmt"
)

func myFunc(firstName string, lastName string) (string) {
	fullName := firstName + " " + lastName
	return fullName
}

func addOne(first int) func() int {
	x := first
	return func() int {
		result := x
		x += 1
		return result
	}
}

func main() {
	result := addOne(0)
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
}