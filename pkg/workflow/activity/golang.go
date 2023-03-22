package activity

import "fmt"

type Golang struct {
	name    string
	handler func(interface{}) interface{}
	input   interface{}
}

func NewGolang(name string, handler func(interface{}) interface{}, input interface{}) *Golang {
	return &Golang{
		name:    name,
		handler: handler,
		input:   input,
	}
}

func (g *Golang) GetName() string {
	return g.name
}

func (g *Golang) Execute() error {
	fmt.Println("Executing Activity Golang", g.name)
	fmt.Println("the output is:", g.handler(g.input))

	return nil
}
