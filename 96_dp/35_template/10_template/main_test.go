package main

import (
	"fmt"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	showActors(NewWomanActor("妈妈"), NewManActor("爸爸"), NewChildActor("儿子"))
}

// showActors 显示演员的装扮信息
func showActors(actors ...IActor) {
	for _, actor := range actors {
		fmt.Println(actor.DressUp())
	}
}
