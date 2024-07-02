package main

import (
	"fmt"
	"todo-cli-golang/todo"
)

const filename = "todo.json"

func main() {
	items, err := todo.LoadItems(filename)
	if err != nil {
		fmt.Println("Error loading items:", err)
		return
	}

	newItem := todo.NewItem(len(items) + 1)
	items = append(items, *newItem)

	err = todo.SaveItems(filename, items)
	if err != nil {
		fmt.Println("Error saving items:", err)
		return
	}
}
