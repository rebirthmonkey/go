package workflow

import "fmt"

type Gateway struct {
	name string
}

func NewGateway(name string) *Gateway {
	return &Gateway{
		name: name,
	}
}

func (g *Gateway) GetName() string {
	return g.name
}

func (g *Gateway) Execute() error {
	fmt.Println("Executing Gateway ", g.name)
	return nil
}
