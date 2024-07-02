package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func NewItem(id int) *Item {
	var task string
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a task you want to add: ")
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading task:", err)
		return nil
	}

	task = strings.TrimSpace(task)

	return &Item{
		ID:   id,
		Task: task,
		Done: false,
	}
}
