package main

import (
	"encoding/json"
	"fmt"
)

type Bird2 struct {
	Species string `json:"birdType"`
	Description string `json:"what it does"`
}

func main(){
	birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`

	var bird Bird2

	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
}
