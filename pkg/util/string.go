package util

import (
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/spf13/cast"
)

func StringContainsAny(str string, substrs []string) bool {
	for _, substr := range substrs {
		if strings.Index(str, substr) > -1 {
			return true
		}
	}
	return false
}

func IsAZaz09(str string, len uint8) bool {
	regex := fmt.Sprintf("^[A-Za-z0-9]{%d}$", len)
	reg, _ := regexp2.Compile(regex, 0)
	match, _ := reg.MatchString(str)
	return match
}

func ToString(obj interface{}) string {
	if obj == nil {
		return ""
	}
	switch o := obj.(type) {
	case *string:
		return *o
	default:
		return cast.ToString(obj)
	}
}
