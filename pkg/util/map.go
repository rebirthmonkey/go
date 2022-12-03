package util

import "strings"

func UniqValMap(keys []string, value string) map[string]string {
	rets := make(map[string]string)
	for _, key := range keys {
		rets[key] = value
	}
	return rets
}

func CopyStringMap(src map[string]string) map[string]string {
	dest := make(map[string]string)
	for k, v := range src {
		dest[k] = v
	}
	return dest
}

func CopyMap(src map[string]interface{}) map[string]interface{} {
	dest := make(map[string]interface{})
	for k, v := range src {
		dest[k] = v
	}
	return dest
}

func DotGet(input map[string]interface{}, name string) interface{} {
	ns := strings.Split(name, ".")
	if !(len(ns) > 0 && ns[0] != "") {
		return input
	}
	r, ok := input[ns[0]]
	if !ok {
		return nil
	}
	ii, ok := r.(map[string]interface{})
	if !ok {
		if len(ns) == 1 {
			return r
		}
		return nil
	}
	return DotGet(ii, strings.Join(ns[1:], "."))
}
