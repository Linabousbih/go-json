package main

import (
	"encoding/json"
	"fmt"

	//"io"
	"log"
)

type Todo struct {
	UserID    int    `json:"-"`
	ID        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {
	todoItem := &Todo{1, 1, "", false}

	todo, err := json.MarshalIndent(todoItem, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(todo))
}
