package calc

import "errors"

func Sub(num1, num2 int) int {
	return num1 - num2
}

func SortArray(arr []int) []int {
	i := 1
	for i < len(arr) {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			aux := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = aux
			j--
		}
		i++
	}
	return arr
}

func Divide(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("o denominador não pode ser 0")
	}
	return num / den, nil
}
