package main

import (
	"encoding/json"
	"fmt"

	//"io"
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

		todoItem := Todo{}

		decoder := json.NewDecoder(response.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&todoItem); err != nil {

			log.Fatal("Decode error:", err)
		}

		fmt.Printf("Data from API: %+v", todoItem)
	}
}
