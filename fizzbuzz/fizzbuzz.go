// Package fizzbuzz allows a program to calculate the answer to FizzBuzz when given any integer
package fizzbuzz

import "strconv"

// Calculate retusn a string representation of the answer to FizzBuzz given an integer value
func Calculate(x int) string {

	var multipleOfThree float64 = float64(x) / 3
	var multipleOfFive float64 = float64(x) / 5

	if (multipleOfThree == float64(int64(multipleOfThree))) && (multipleOfFive == float64(int64(multipleOfFive))) {
		return "FizBuzz"
	}

	if multipleOfThree == float64(int64(multipleOfThree)) {
		return "Fizz"
	}

	if multipleOfFive == float64(int64(multipleOfFive)) {
		return "Buzz"
	}

	return strconv.Itoa(x)
}
