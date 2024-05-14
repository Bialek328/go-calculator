package main

import (
	"errors"
	"fmt"

	"example.com/calculations"
)

func getOperator() string {
	fmt.Println("Operator: ")
	input := ""
	fmt.Scanln(&input)
	return input
}

func getNumbers() (float32, float32) {
	var number1, number2 float32
	fmt.Println("First number:")
	fmt.Scanln(&number1)
	fmt.Println("Second number:")
	fmt.Scanln(&number2)
	return number1, number2
}

func handleCalculation(opr string, number1, number2 float32) (float32, error) {
	switch opr {
	case "+":
		return calculations.Add(number1, number2), nil

	case "-":
		return calculations.Substract(number1, number2), nil

	case "*":
		return calculations.Multiply(number1, number2), nil

	case "/":
		return calculations.Divide(number1, number2)

	default:
		return 0, errors.New("invalid operator")
	}
}

func main() {
	opr := getOperator()
	// opr, err := validateOperator(str)
	number1, number2 := getNumbers()
	result, err := handleCalculation(opr, number1, number2)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %f\n", result)
	}
}
