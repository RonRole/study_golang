package main

import (
	"fmt"
)

func main() {
	var myint float64
	for i := 0; i < 10; i++ {
		myint += 0.1
	}
	fmt.Println(myint)
}

//float 