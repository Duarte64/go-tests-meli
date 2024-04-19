package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubNegative(t *testing.T) {
	num1 := 3
	num2 := 5
	expectedResult := -2

	result := Sub(num1, num2)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestSubPositivo(t *testing.T) {
	num1 := 5
	num2 := 2
	expectedResult := 3

	result := Sub(num1, num2)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestSortArray(t *testing.T) {
	unsortedSlice := []int{5, 2, 8, 3, 1}
	expectedSortedSlice := []int{1, 2, 3, 5, 8}

	sortedSlice := SortArray(unsortedSlice)

	assert.Equal(t, expectedSortedSlice, sortedSlice, "Os slices devem ser iguais após a ordenação")
}

func TestDivide(t *testing.T) {
	num := 8
	den := 4

	expectedResult := 2

	result, _ := Divide(num, den)

	assert.Equal(t, expectedResult, result, "8/4 deve ser igual a 2")
}

func TestDivideByZero(t *testing.T) {
	num := 7
	den := 0

	_, err := Divide(num, den)

	assert.NotNil(t, err)
}
