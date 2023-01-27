package fizzbuzz

import "strconv"

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
