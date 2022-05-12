package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User3 struct {
	Name          string    `json:"name"`
	Password      string    `json:"-"`                       // 设置为private项
	PreferredFish []string  `json:"preferredFish,omitempty"` // 忽略为null的项
	CreatedAt     time.Time `json:"createdAt"`
}

func main() {
	u := &User3{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
