package main

import (
	"fmt"

	"github.com/go-micah/fizzbuzz/fizzbuzz"
)

func main() {
	for x := 1; x <= 100; x++ {
		fmt.Println(fizzbuzz.Calculate(x))
	}
}
