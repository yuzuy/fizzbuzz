package fizzbuzz

import "fmt"

func FizzBuzz(i int64) string {
	if i%15 == 0 {
		return "FizzBuzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	if i%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(i)
}
