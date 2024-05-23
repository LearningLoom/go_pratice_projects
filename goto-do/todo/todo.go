package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

const todoFile = "./todos.json"

type Todo struct {
	Id          string    `json:"id"`
	Task        string    `json:"task"`
	Description string    `json:"description"`
	CreatedAT   time.Time `json:"created_at"`
}

func New(task, description string) ([]*Todo, error) {
	isEmpty, err := isJSONFileEmpty(todoFile)
	if err != nil {

		return nil, err
	}
	if isEmpty {
		createJsonEmptyArrayFile(todoFile)
	}
	existingTodos, err := Get()
	if err != nil {
		return nil, err
	}

	// create a new todo
	todo := &Todo{
		Id:          randomString(),
		CreatedAT:   time.Now(),
		Task:        task,
		Description: description,
	}
	todos := append(existingTodos, todo)
	if err := WriteTodoToFile(todos); err != nil {
		return nil, err
	}

	return todos, nil

}

func (todo *Todo) Save(id, task, description string) *Todo {
	todo.Task = task
	todo.Description = description
	return todo
}

func WriteTodoToFile(todos []*Todo) error {
	jsonData, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	err = os.WriteFile(todoFile, jsonData, 0775)
	if err != nil {
		return err
	}
	return nil
}

func Get() ([]*Todo, error) {
	isEmpty, err := isJSONFileEmpty(todoFile)
	if err != nil {

		return nil, err
	}
	if isEmpty {
		err = errors.New("there are no Todos, Please create some")
		return nil, err
	}
	data, err := os.ReadFile(todoFile)

	if err != nil {
		return nil, err
	}
	var todo []*Todo

	err = json.Unmarshal(data, &todo)

	if err != nil {
		return nil, err
	}
	return todo, nil
}
