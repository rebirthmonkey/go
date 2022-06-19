package main

import "fmt"

type testImpl struct {
	name string
}

func (t testImpl) create() string {
	fmt.Println("testImpl create")
	return t.name
}

func (t testImpl) get() string {
	fmt.Println("testImpl get")
	return t.name
}

func (t testImpl) delete() string {
	fmt.Println("testImpl delete")
	return t.name
}
