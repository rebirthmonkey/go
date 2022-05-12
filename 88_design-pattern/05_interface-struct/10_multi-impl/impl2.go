package main

import "fmt"

type testImpl2 struct {
	name string
}

func (t testImpl2) create() string {
	fmt.Println("testImpl2 create")
	return t.name
}

func (t testImpl2) get() string {
	fmt.Println("testImpl2 get")
	return t.name
}

var test2 = testImpl2{"test2"}

func getTest2() TestInterface {
	return test2
}
