package mgr

import (
	"fmt"
)

type Manager struct {
	Plugins map[string]string
}

func (s *Manager) Show() {
	for key, val := range s.Plugins {
		fmt.Println("key: ", key, " value: ", val)
	}
}
