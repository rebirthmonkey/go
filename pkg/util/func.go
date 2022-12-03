package util

import (
	"reflect"
	"runtime"
	"strings"
)

func GetFuncName(fn interface{}) string {
	fnPath := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	fnPaths := strings.Split(fnPath, "/")
	fnPName := fnPaths[len(fnPaths)-1]
	dotIdx := strings.Index(fnPName, ".")
	if dotIdx == -1 {
		return fnPName
	} else {
		return fnPName[dotIdx+1:]
	}
}
