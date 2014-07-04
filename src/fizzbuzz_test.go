package fizzbuzz

import (

	"github.com/stretchr/testify/assert"
	"testing"
	"fizzbuzz"

)

func TestVerifcaUno(t *testing.T) {

	assert.Equal(t,  Fizzbuzz(1), "1")
}

func TestVerifcaDos(t *testing.T) {

	assert.Equal(t, Fizzbuzz(2), "2")
}

func TestVerificaFizz(t *testing.T) {

	assert.Equal(t, Fizzbuzz(3), "Fizz")
}

func TestVerificaContieneTres(t *testing.T) {

	assert.Equal(t, Fizzbuzz(13), "Fizz")
}

func testVerificaBuzz(t *testing.T) {

	assert.Equal(t, Fizzbuzz(5), "Buzz")
}

func TestVerificaContieneCinco(t *testing.T) {

	assert.Equal(t, Fizzbuzz(52), "Buzz")
}

func TestVerificaFizzBuzz(t *testing.T) {

	assert.Equal(t, Fizzbuzz(15), "FizzBuzz")
}
