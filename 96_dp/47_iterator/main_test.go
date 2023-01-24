package main

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	class := NewClass("三年级一班", "王明", "数学课")
	class.AddStudent(NewStudent("张三", 389),
		NewStudent("李四", 378),
		NewStudent("王五", 347))

	fmt.Printf("%s成员如下:\n", class.Name())
	classIterator := class.CreateIterator()
	for classIterator.HasMore() {
		member := classIterator.Next()
		fmt.Println(member.Desc())
	}
}
