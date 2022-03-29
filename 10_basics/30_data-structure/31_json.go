package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species string
	Description string
}

func main(){
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`

	var bird Bird

	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
}
