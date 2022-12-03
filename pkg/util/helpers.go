package util

import (
	"math"
	"sync"
)

// Paginate
func Paginate(total, pageSize int, onPage func(offset, limit int) bool) {
	for i := 0; i < total; i += pageSize {
		if i+pageSize > total {
			pageSize = total - i
		}
		if !onPage(i, pageSize) {
			return
		}
	}
}

// concurrent version of Paginate
func PaginateC(total, pageSize int, onPage func(offset, limit int)) {
	var wg sync.WaitGroup
	Paginate(total, pageSize, func(offset, limit int) bool {
		wg.Add(1)
		go func() {
			onPage(offset, limit)
			wg.Done()
		}()
		return true
	})
	wg.Wait()
}

func PaginateV(total, maxC int, onPage func(offset, limit int)) {
	PaginateC(total, int(math.Ceil(float64(total)/float64(maxC))), onPage)
}

// maxC max concurrence
// maxP max pageSize
func PaginateCV(total, maxC, maxP int, onPage func(offset, limit int)) {
	if total <= maxC*maxP {
		PaginateV(total, maxC, onPage)
		return
	}
	Paginate(total, maxC*maxP, func(subOffset, subTotal int) bool {
		PaginateV(subTotal, maxC, func(offset, limit int) {
			onPage(subOffset+offset, limit)
		})
		return true
	})
}
