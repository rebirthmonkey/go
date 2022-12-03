package util

import (
	"testing"
)

func Test_Paginate(t *testing.T) {
	offsets := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	page := 0
	Paginate(100, 10, func(offset, limit int) bool {
		if page >= len(offsets) {
			t.Fatal("paginate total page error")
		}
		if offset != offsets[page] {
			t.Fatal("offset error")
		}
		page++
		return true
	})
}
