package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	// intialize i, which we use for our boolean expression in the for loop
	i := 1000

	// execute a loop while i > 100
	for i > 100 {
		// get a random number between 1 and 1001
		i = rand.Intn(1000) + 1
		fmt.Println("i is", i)
		if i > 100 {
			fmt.Println("i is", i, "so loop keeps going")
		}
	}

	fmt.Println("Got", i, "and broke out of loop")
}
