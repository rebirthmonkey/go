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
	/*
	{
	  "species": "pigeon",
	  "decription": "likes to perch on rocks"
	}
	 */
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`

	var bird Bird

	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
}
