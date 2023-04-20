package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User2 struct {
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	PreferredFish []string  `json:"preferredFish"`
	CreatedAt     time.Time `json:"createdAt"`
}

func main() {
	u := &User2{
		Name:      "Wukong",
		Password:  "Password",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "user", "-------")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
