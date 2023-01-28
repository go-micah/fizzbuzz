/*
Fizzy is a sample command line program to test the [github.com/go-micah/fizzbuzz] package.

Usage:

	./bin/fizzy
*/
package main

import (
	"fmt"

	"github.com/go-micah/fizzbuzz"
)

func main() {
	for x := 1; x <= 100; x++ {
		fmt.Println(fizzbuzz.Calculate(x))
	}
}
