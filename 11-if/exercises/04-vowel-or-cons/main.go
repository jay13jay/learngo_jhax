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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 2 || len(os.Args[1]) != 1 {
		fmt.Println("Give me a letter")
		os.Exit(0)
	}
	let := os.Args[1]
	if strings.IndexAny(let, "aeiou") >= 0 {
		fmt.Printf("%q is a vowel\n", let)
	} else if strings.IndexAny(let, "yw") >= 0 {
		fmt.Printf("%q is sometimes a vowel, sometimes not\n", let)
	} else {
		// in reality this can catch far more than consonants. you could use index any and pass in the rest of the alphabet
		// but that would be really ugly and beyond scope
		fmt.Printf("%q is a consonant\n", let)
	}
}
