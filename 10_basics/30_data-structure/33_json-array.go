package main

import (
	"encoding/json"
	"fmt"
)

type Bird3 struct {
	Species string
	Description string
}

func main(){
	birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},
				  {"species":"eagle","description":"bird of prey"}]`

	var birds []Bird3

	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("Birds : %+v", birds)
}
