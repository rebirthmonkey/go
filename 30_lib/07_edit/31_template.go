package main

import (
	"os"
	"text/template"
)

func main() {
	text := "call: {{ call .x .y .z }} \n"
	tmpl, _ := template.New("test").Parse(text)
	tmpl.Execute(os.Stdout, map[string]interface{}{
		"x": func(x, y int) int { return x+y},
		"y": 2,
		"z": 3,
	})
}


