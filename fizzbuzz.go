package fizzbuzz

import "strconv"

func FizzBuzz(value int) string {
	if value < 1 {
		return ""
	}

	if value%15 == 0 {
		return "FizzBuzz"
	}
	if value%5 == 0 {
		return "Buzz"
	}
	if value%3 == 0 {
		return "Fizz"
	}

	return strconv.Itoa(value)
}
