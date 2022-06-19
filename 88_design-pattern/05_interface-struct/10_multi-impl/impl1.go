package main

import "fmt"

type testImpl1 struct {
	name string
}

func (t testImpl1) create() string {
	fmt.Println("testImpl1 create")
	return t.name
}

func (t testImpl1) get() string {
	fmt.Println("testImpl1 get")
	return t.name
}

var test1 = testImpl1{"test1"}

func getTest1() TestInterface {
	return test1
}
