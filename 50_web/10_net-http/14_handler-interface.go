package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page3 struct {
	Title string
	Body  []byte
}

func loadPage3(title string) (*Page3, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page3{Title: title, Body: body}, nil
}

type viewHandler3 struct{}

func (viewHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage3(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.Handle("/view/", viewHandler3{})
	http.ListenAndServe(":8080", nil)
}
