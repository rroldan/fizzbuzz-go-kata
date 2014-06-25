package fizzbuzz

import "strconv"
import "strings"




func Fizzbuzz(inputValue int) string {

	const (
		FIZZ      string = "Fizz"
		BUZZ      string = "Buzz"
		FIZZ_BUZZ string = "FizzBuzz"
	)

	var (
		inputValueStr string
	)

	inputValueStr = strconv.Itoa(inputValue)

	if (inputValue%3 == 0) && (inputValue%5 == 0) {
		return FIZZ_BUZZ
	}

	if (inputValue%3 == 0) || strings.Contains(inputValueStr, "3") {
		return FIZZ
	}

	if (inputValue%5 == 0) || strings.Contains(inputValueStr, "5") {
		return BUZZ
	}

	return inputValueStr
}

