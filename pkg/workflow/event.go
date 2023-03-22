package workflow

import "fmt"

type Event struct {
	name string
}

func NewEvent(name string) *Event {
	return &Event{
		name: name,
	}
}

func (e *Event) GetName() string {
	return e.name
}

func (e *Event) Execute() error {
	fmt.Println("Executing Event ", e.name)
	return nil
}
