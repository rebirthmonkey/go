package main

import (
	"bytes"
	"fmt"
	"io"
)

// Paper 纸张，包含读取内容的方法，拷贝纸张的方法，作为原型模式接口
type Paper interface {
	io.Reader
	Clone() Paper
}

// Newspaper 报纸 实现原型接口
type Newspaper struct {
	headline string
	content  string
}

func NewNewspaper(headline string, content string) *Newspaper {
	return &Newspaper{
		headline: headline,
		content:  content,
	}
}

func (np *Newspaper) Read(p []byte) (n int, err error) {
	buf := bytes.NewBufferString(fmt.Sprintf("headline:%s,content:%s", np.headline, np.content))
	return buf.Read(p)
}

func (np *Newspaper) Clone() Paper {
	return &Newspaper{
		headline: np.headline + "_copied",
		content:  np.content,
	}
}

// Resume 简历 实现原型接口
type Resume struct {
	name       string
	age        int
	experience string
}

func NewResume(name string, age int, experience string) *Resume {
	return &Resume{
		name:       name,
		age:        age,
		experience: experience,
	}
}

func (r *Resume) Read(p []byte) (n int, err error) {
	buf := bytes.NewBufferString(fmt.Sprintf("name:%s,age:%d,experience:%s", r.name, r.age, r.experience))
	return buf.Read(p)
}

func (r *Resume) Clone() Paper {
	return &Resume{
		name:       r.name + "_copied",
		age:        r.age,
		experience: r.experience,
	}
}
