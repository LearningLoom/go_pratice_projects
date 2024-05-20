package main

import (
	"errors"
	"fmt"
)

// * this is a basic go project here we will learn how to handle user input and how to pass values to functions
// * furthermore we will also learn about inifnite loops here

//  ! THIS WILL ACCEPT 2 NUMBERS AND PERFORM ARTHMATIC OPERATIONS IN IT

// * This is the main function that wil execute all the other functions.
func main() {
	for {
		value1, value2, operator, err := getUserValues()

		if err != nil {
			fmt.Println(err)
		}
		switch operator {
		case "+":
			result := value1 + value2
			fmt.Printf("The result of %2.f + %2.f is:  %2.f \n", value1, value2, result)
		case "-":
			result := value1 - value2
			fmt.Printf("The result of %2.f - %2.f is:  %2.f \n", value1, value2, result)
		case "*":
			result := value1 * value2
			fmt.Printf("The result of %2.f * %2.f is:  %2.f \n", value1, value2, result)
		case "/":
			result := value1 / value2
			fmt.Printf("The result of %2.f / %2.f is:  %2.f \n", value1, value2, result)
		default:
			fmt.Println("Please provide a valid operator, within the given range")

		}

	}

}

// * function to get values from the user
func getUserValues() (float64, float64, string, error) {
	var value1 float64
	var value2 float64
	var operator string
	fmt.Println("Basic Calculator ")
	fmt.Println("____________________________________________")
	fmt.Println("Please enter the first Number: ")
	fmt.Scan(&value1)
	fmt.Println("Please enter the second number: ")
	fmt.Scan(&value2)
	fmt.Println("Please enter the operation that you want to do: +, -, * or / : ")
	fmt.Scan(&operator)

	if value1 == 0.0 || value2 == 0.0 || operator == "" {
		err := errors.New("please enter valid data")
		return 0.0, 0.0, "", err
	}
	return value1, value2, operator, nil
}
