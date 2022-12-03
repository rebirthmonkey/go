package util

import (
	"testing"
)

func TestMap_Get(t *testing.T) {
	param := map[string]interface{}{
		"path1": map[string]interface{}{
			"path2": map[string]interface{}{
				"path3": "path3",
			},
		},
	}
	cases := []string{
		"path1",
		"path1.path2",
		"path1.path2.path3",
		"path1.path2.path3.path4",
		"path2"}
	for i, p := range cases {
		r := DotGet(param, p)
		switch {
		case i < 2:
			if _, ok := r.(map[string]interface{}); !ok {
				t.Fatal("should got map[string]interface{}")
			}
		case i == 2:
			if _, ok := r.(string); !ok {
				t.Fatal("should got string")
			}
		default:
			if r != nil {
				t.Fatal("should got nil")
			}
		}
	}
}
