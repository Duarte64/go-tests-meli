package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciZero(t *testing.T) {
	num := 0
	expectedResult := 0

	result := Fibonacci(num)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestFibonacciUm(t *testing.T) {
	num := 1
	expectedResult := 1

	result := Fibonacci(num)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestFibonacciMaior(t *testing.T) {
	num := 30
	expectedResult := 832040

	result := Fibonacci(num)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}
