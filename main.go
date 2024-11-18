package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todoo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

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

		data := string(bodyBytes)
		fmt.Println(data)
	}
}
