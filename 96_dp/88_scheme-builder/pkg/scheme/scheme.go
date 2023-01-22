package scheme

import (
	"fmt"
)

type Scheme struct {
	Plugins map[string]string
}

func (s *Scheme) Show() {
	for key, val := range s.Plugins {
		fmt.Println("key: ", key, " value: ", val)
	}
}
