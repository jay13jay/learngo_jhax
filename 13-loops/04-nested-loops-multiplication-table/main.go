// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// EXERCISE: Get the `max` from the user
//           And create the table according to that.

const (
	errArgs = "Please provide a single integer argument"
	errType = "Argument not an integer"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(errArgs)
		return
	}
	max, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(errType)
		return
	}

	// print the header
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
