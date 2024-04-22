package calculations

import "errors"

func Add(x, y float32) float32 {
	sum := x + y
	return sum
}

func Substract(x, y float32) float32 {
	result := x - y
	return result
}

func Multiply(x, y float32) float32 {
	result := x * y
	return result
}

func Divide(x, y float32) (float32, error) {
	if y == 0 {
		return -1, errors.New("0 division error")
	}
	result := x / y
	return result, nil
}
