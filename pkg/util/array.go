package util

import (
	"errors"
	"sync"

	"github.com/thoas/go-funk"
)

// ArrayChunkThen split array in small piece then handle them
func ArrayChunkThen(arr []string, chunkSize int, then func(chunk []string, xth int) error) error {
	if len(arr) == 0 {
		return nil
	}
	total := len(arr)
	start, limit := 0, chunkSize
	idx := 0
	for {
		end := start + limit
		if end > total {
			end = total
			limit = end - start
		}
		if err := then(arr[start:end], idx); err != nil {
			return err
		}
		start = end
		if start >= total {
			break
		}
		idx += 1
	}
	return nil
}

func XthSlice(input []string, chunkSize, xth int) []string {
	start := xth * chunkSize
	end := (xth + 1) * chunkSize
	if end > len(input) {
		end = len(input) - 1
	}
	if start > end {
		return []string{}
	}
	return input[start:end]
}

// ArrayChunkThenV a variant of ArrayChunkThen, this may be the one that most of us wanted
func ArrayChunkThenV(arr []string, chunkSize int, then func(chunk []string, xth int) bool) {
	ArrayChunkThen(arr, chunkSize, func(chunk []string, xth int) error {
		if !then(chunk, xth) {
			return errors.New("I'm out!")
		}
		return nil
	})
}

// ArrayChunkThenC concurrence version of ArrayChunkThen
func ArrayChunkThenC(arr []string, chunkSize int, then func(chunk []string, xth int)) {
	var wg sync.WaitGroup
	ArrayChunkThen(arr, chunkSize, func(chunk []string, xth int) error {
		wg.Add(1)
		go func() {
			then(chunk, xth)
			wg.Done()
		}()
		return nil
	})
	wg.Wait()
}

// ArrayEqual check if arr1 and arr2 are equal without order
func ArrayEqual(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	diff := make(map[string]int, len(arr1))
	for _, val := range arr1 {
		diff[val]++
	}
	for _, val := range arr2 {
		if _, ok := diff[val]; !ok {
			return false
		}
		diff[val] -= 1
		if diff[val] == 0 {
			delete(diff, val)
		}
	}
	return len(diff) == 0
}

// ArrayContain check if arr1 contain all the elements of arr2
func ArrayContain(arr1 []string, arr2 []string) bool {
	if len(arr1) < len(arr2) {
		return false
	}
	diff := make(map[string]bool, len(arr1))
	for _, val := range arr1 {
		diff[val] = true
	}
	for _, val := range arr2 {
		if _, ok := diff[val]; !ok {
			return false
		}
	}
	return true
}

// ArrayContainAny check if arr1 contain any elements of arr2
func ArrayContainAny(arr1 []string, arr2 []string) bool {
	return len(funk.IntersectString(arr1, arr2)) > 0
}

func ArrayHas(arr []string, val string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}
