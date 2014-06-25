package fizzbuzz

import (

	"github.com/stretchr/testify/assert"
	"testing"
	"fizzbuzz"

)

func TestVerifcaUno(t *testing.T) {

	assert.Equal(t,  fizzbuzz.Fizzbuzz(1), "1")
}

func TestVerifcaDos(t *testing.T) {

	assert.Equal(t, fizzbuzz.Fizzbuzz(2), "2")
}

func TestVerificaFizz(t *testing.T) {

	assert.Equal(t, fizzbuzz.Fizzbuzz(3), "Fizz")
}

func TestVerificaContieneTres(t *testing.T) {

	assert.Equal(t, fizzbuzz.Fizzbuzz(13), "Fizz")
}

func testVerificaBuzz(t *testing.T) {

	assert.Equal(t, fizzbuzz.Fizzbuzz(5), "Buzz")
}

func TestVerificaContieneCinco(t *testing.T) {

	assert.Equal(t, fizzbuzz.Fizzbuzz(52), "Buzz")
}

func TestVerificaFizzBuzz(t *testing.T) {

	assert.Equal(t, fizzbuzz.Fizzbuzz(15), "FizzBuzz")
}
