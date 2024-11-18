package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	//This will get executed later when the main function ends to prevent memory leaks
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		todoItem := Todo{}

		json.Unmarshal(bodyBytes, &todoItem)

		fmt.Printf("Data from API: %+v", todoItem)
	}
}
