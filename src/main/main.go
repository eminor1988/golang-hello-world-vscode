package main

import (
	"fmt"
	"hello"
)

func main() {
	const (
		firstName string = "Arthur"
		lastName  string = "Schopenhauer"
	)

	helloMessage := hello.HelloMessage()
	for i := 0; i < 2; i++ {
		fmt.Println(i, helloMessage, firstName, lastName)
	}
}
