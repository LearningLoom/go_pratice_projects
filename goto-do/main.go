package main

import (
	"bufio"
	"errors"
	"fmt"
	"go_to-do/todo"
	"os"
	"strings"
)

func main() {
	todos, err := inputSelector()
	if err != nil {
		fmt.Println(err)
	}
	for _, todo := range todos {
		fmt.Printf("ID: %s, Task: %s, Description: %s\n", todo.Id, todo.Task, todo.Description)
	}
}

func getInitialUserInput() int {
	fmt.Println("-----------------------------------------")
	fmt.Println(strings.ToUpper("welcome to Todo app"))
	fmt.Println("Please choose your options below")
	fmt.Println("++++++++++++++++++++++")
	fmt.Print("1. for gettting all todos.\n2. Creating a Todo.\n")
	var option int
	fmt.Scan(&option)
	return option
}

func inputSelector() ([]*todo.Todo, error) {
	option := getInitialUserInput()
	switch option {
	case 1:
		todos, err := todo.Get()
		if err != nil {
			return nil, err
		}
		return todos, nil
	case 2:
		title, description, err := collectTodoData()
		if err != nil {
			return nil, err
		}
		newTodo, err := todo.New(title, description)
		if err != nil {
			return nil, err
		}
		return newTodo, nil
	default:
		err := errors.New("please enter a vlid input")
		return nil, err
	}

}

func collectTodoData() (string, string, error) {
	fmt.Println("Please the Title for the Todo: ")
	// as fmt.scan() cannot accept input with spaces and multiline input we will use another way to do it .
	title, err := readCliInput()

	if err != nil {
		return "", "", err
	}
	fmt.Println("Please enter the description: ")
	description, err := readCliInput()

	fmt.Print("++++++++++++++++++++++++++++++++++++\n")

	if err != nil {
		return "", "", err
	}
	return title, description, nil

}

func readCliInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text, nil
}
