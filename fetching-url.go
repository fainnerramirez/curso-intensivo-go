package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DogResponseDTO struct {
	Message string `json:message`
	Status  string `json:status`
}

func main() {

	response, err := http.Get("https://dog.ceo/api/breeds/image/random")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic("Error en la petición HTTP")
	}

	var dog DogResponseDTO

	err = json.NewDecoder(response.Body).Decode(&dog)

	if err != nil {
		panic(err)
	}

	fmt.Println(dog.Message)
	fmt.Print(dog.Status)
}
