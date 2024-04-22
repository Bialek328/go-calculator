package main

import (
	"errors"
	"fmt"
)

var operators = []string{"+", "-", "*", "/"}

func validateOperator(input string) (string, error) {
	for _, opr := range operators {
		if input == opr {
			return opr, nil
		}
	}
	return "", errors.New("invalid operator!")
}

func getUserInput() string {
	fmt.Println("Input: ")
	input := ""
	fmt.Scanln(&input)
	return input
}

func main() {
	str := getUserInput()
	opr, err := validateOperator(str)
	if err != nil {
		fmt.Printf("Error: ", err)
	}
	if 
}
