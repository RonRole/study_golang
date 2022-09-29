package main

import (
	"fmt"
)

func main() {
	type Person struct {
		name string
		age	 int
	}

	type Team struct {
		name string
		players [2]Person
	}

	players := [...]Person {
		Person {
			name: "Forrest",
			age: 18,
		},
		Person {
			name: "Gordon",
			age: 24,
		},
	}

	celtic := Team {
		name: "Celtic FC",
		players: players,
	}

	fmt.Println(celtic)
}