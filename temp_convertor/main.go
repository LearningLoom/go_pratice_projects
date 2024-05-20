package main

import (
	"fmt"
	"strings"

	"github.com/syscll/tempconv"
)

// ! in this programe we will make use of a third party liberary and we will learn how to use it
//  ! Here are some steps that you need to do before writing code 1. " go mod init <project_name" 1. create a main.go file
//  ! the liberary that we will use is github.com/syscll/tempconv

func main() {
	option, temp := getValues()
	value := convertTo(option, temp)
	fmt.Print(value)
}

func getValues() (int, float64) {
	fmt.Println(strings.ToUpper("welcome to Temperature convertor !!!"))
	fmt.Println(" How do you want to convert ? : ")
	fmt.Print("1. Celcius ->  Fahrenheit\n2. Fahrenheit ->  Celcius\n3. Celcius ->  Kelvin\n4. Fahrenheit ->  Kelvin\n5. Kelvin ->  Celcius\n6. Kelvin -> Fahrenheit\nChoose by selecting the number i'e 1,2,3.4,5 or 6 \n")
	var option int
	fmt.Scan(&option)
	fmt.Println("Please enter the temperature: ")
	var temp float64
	fmt.Scan(&temp)
	return option, temp
}

func convertTo(option int, temp float64) string {
	var convertedTemp float64
	var result string
	switch option {
	case 1:
		convertedTemp = float64(tempconv.CelsiusToFahrenheit(tempconv.Celsius(temp)))
		result = fmt.Sprintf("Fahrenheit value is %2.f F \n", convertedTemp)
	case 2:
		convertedTemp = float64(tempconv.FahrenheitToCelsius(tempconv.Fahrenheit(temp)))
		result = fmt.Sprintf("Fahrenheit value is %2.f C \n", convertedTemp)
	case 3:
		convertedTemp = float64(tempconv.CelsiusToKelvin(tempconv.Celsius(temp)))
		result = fmt.Sprintf("Fahrenheit value is %2.f K \n", convertedTemp)
	case 4:
		convertedTemp = float64(tempconv.FahrenheitToKelvin(tempconv.Fahrenheit(temp)))
		result = fmt.Sprintf("Fahrenheit value is %2.f K \n", convertedTemp)
	case 5:
		convertedTemp = float64(tempconv.KelvinToCelcius(tempconv.Kelvin(temp)))
		result = fmt.Sprintf("Fahrenheit value is %2.f C \n", convertedTemp)
	case 6:
		convertedTemp = float64(tempconv.KelvinToFahrenheit(tempconv.Kelvin(temp)))
		result = fmt.Sprintf("Fahrenheit value is %2.f F \n", convertedTemp)
	default:
		result = fmt.Sprintln("Please choose a valid Option ")
	}
	return result
}
