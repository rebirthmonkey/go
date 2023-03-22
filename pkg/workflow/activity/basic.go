package activity

import "fmt"

type Basic struct {
	name string
}

func NewBasic(name string) *Basic {
	return &Basic{
		name: name,
	}
}

func (b *Basic) GetName() string {
	return b.name
}

func (b *Basic) Execute() error {
	fmt.Println("Executing Activity Basic", b.name)

	return nil
}
