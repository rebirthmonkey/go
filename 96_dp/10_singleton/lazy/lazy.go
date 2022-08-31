package lazy

import (
	"sync"
)

type singleton struct {
	mode string
}

var ins *singleton
var once sync.Once

func GetInsOr() *singleton {
	once.Do(func() {
		ins = &singleton{"off"}
	})
	return ins
}