package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	copier := NewCopier("云打印机")
	oneNewspaper := NewNewspaper("Go是最好的编程语言", "Go语言十大优势")
	oneResume := NewResume("小明", 29, "5年码农")

	otherNewspaper := copier.copy(oneNewspaper)
	copyNewspaperMsg := make([]byte, 100)
	byteSize, _ := otherNewspaper.Read(copyNewspaperMsg)
	fmt.Println("copyNewspaperMsg:" + string(copyNewspaperMsg[:byteSize]))

	otherResume := copier.copy(oneResume)
	copyResumeMsg := make([]byte, 100)
	byteSize, _ = otherResume.Read(copyResumeMsg)
	fmt.Println("copyResumeMsg:" + string(copyResumeMsg[:byteSize]))
}

// Copier 复印机
type Copier struct {
	name string
}

func NewCopier(n string) *Copier {
	return &Copier{name: n}
}

func (c *Copier) copy(paper Paper) Paper {
	fmt.Printf("copier name:%v is copying:%v ", c.name, reflect.TypeOf(paper).String())
	return paper.Clone()
}
