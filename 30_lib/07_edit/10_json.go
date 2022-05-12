package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name          string
	Password      string
	PreferredFish []string
	CreatedAt     time.Time
}

func main() {
	u := &User{
		Name:      "Sammy the Shark", // 如果改为小写，则为private，无法对外输出
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	//out, err := json.MarshalIndent(u, "", "  ")
	out, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("user struct is:", u)
	fmt.Println("user json output is:", string(out))
}
