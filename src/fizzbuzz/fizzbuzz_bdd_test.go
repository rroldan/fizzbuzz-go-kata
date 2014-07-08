package fizzbuzz

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFizzBuzzBDD(t *testing.T) {
	t.Parallel()

	Convey("Given número divisible por 3. ", t, func() {
			x := 3

		Convey("When llama función FizzBuzz", func() {
					val := Fizzbuzz(x)

					Convey("Then retorna valor Fizz", func() {
							So(val, ShouldEqual, "Fizz")

					})

			})
	})

	Convey("Given número divisible por 5. ", t, func() {
			x := 5

			Convey("When llama función FizzBuzz", func() {
			val := Fizzbuzz(x)

				Convey("Then retorna valor Buzz", func() {
						So(val, ShouldEqual, "Buzz")

				})

			})
	})

	Convey("Given número divisible por 3 y por 5. ", t, func() {
			x := 15

			Convey("When llama función FizzBuzz", func() {
			val := Fizzbuzz(x)

				Convey("Then retorna valor FizzBuzz", func() {
					So(val, ShouldEqual, "FizzBuzz")

				})

			})
	})

	Convey("Given Un número divisible por 3 o si incluye un 3 en el número", t, func() {
			x := 53

			Convey("When llama función FizzBuzz", func() {
			val := Fizzbuzz(x)

				Convey("Then retorna valor Fizz", func() {
					So(val, ShouldEqual, "Fizz")

				})

			})
	})

	Convey("Given Un número divisible por 5 o si incluye un 5 en el número", t, func() {
			x := 56

			Convey("When llama función FizzBuzz", func() {
					val := Fizzbuzz(x)

					Convey("Then retorna valor Buzz", func() {
							So(val, ShouldEqual, "Buzz")

					})

			})
	})





}

