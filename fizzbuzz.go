// Package fizzbuzz allows a program to calculate the answer to FizzBuzz when given any integer
package fizzbuzz

import "strconv"

// Calculate returns a string representation of the answer to FizzBuzz given an integer value
func Calculate(x int) string {

	if x == 0 {
		return "0"
	}

	multipleOfThree := x % 3
	multipleOfFive := x % 5

	if (multipleOfThree == 0) && (multipleOfFive == 0) {
		return "FizzBuzz"
	}

	if multipleOfThree == 0 {
		return "Fizz"
	}

	if multipleOfFive == 0 {
		return "Buzz"
	}

	return strconv.Itoa(x)
}
