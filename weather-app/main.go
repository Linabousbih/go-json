package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type myjson struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Visibility int `json:"visibility"`
	Main       struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	BASE_URL := os.Getenv("BASE_URL")
	API_KEY := os.Getenv("API_KEY")

	fmt.Print("Where do you want to check the weather:")
	var city string
	fmt.Scanln(&city)

	searchURL := fmt.Sprintf("%s?appid=%s&q=%s", BASE_URL, API_KEY, city)

	response, err := http.Get(searchURL)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)

		weather := myjson{}

		json.Unmarshal(bytes, &weather)

		fmt.Printf("Today in %v the weather is: %v. The temperature is %2.2v ºC and the humidity is %v", city, weather.Weather[0].Description, weather.Main.Temp-273.15, weather.Main.Humidity)
	}
}
