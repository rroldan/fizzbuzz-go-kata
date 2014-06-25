package main

import "fmt"
import "fizzbuzz"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz.Fizzbuzz(i))
	}
}

