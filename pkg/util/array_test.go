package util

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func Test_ChunkArrayThen(t *testing.T) {
	arr := make([]string, 100)
	expect := 0
	for i := 0; i < 100; i++ {
		arr[i] = strconv.Itoa(i)
		expect += i
	}
	result := 0
	chunkThen := func(chunk []string, xth int) error {
		for _, v := range chunk {
			vi, _ := strconv.Atoi(v)
			result += vi
		}
		return nil
	}
	ArrayChunkThen(arr, 9, chunkThen)
	if result != expect {
		t.Fatalf("ChunkArrayThen not work")
	}
	result = 0
	ArrayChunkThen(arr, 200, chunkThen)
	if result != expect {
		t.Fatalf("ChunkArrayThen not work")
	}
	result = 0
	ArrayChunkThen(arr, 1, chunkThen)
	if result != expect {
		t.Fatalf("ChunkArrayThen not work")
	}
}

func Test_ArrayEqual(t *testing.T) {
	arr1 := make([]string, 10)
	arr2 := make([]string, 10)
	for i := 0; i < 10; i++ {
		val := strconv.Itoa(rand.Intn(100))
		arr1[i] = val
		arr2[i] = val
	}
	if !ArrayEqual(arr1, arr2) {
		t.Fatalf("array equal with order error")
	}
	sortedArr1 := sort.StringSlice(arr1)
	sortedArr1.Sort()

	if !ArrayEqual(sortedArr1, arr2) {
		t.Fatalf("array equal without order error")
	}
}

func Test_ArrayContain(t *testing.T) {
	arr1 := make([]string, 10)
	arr2 := make([]string, 10)
	for i := 0; i < 10; i++ {
		val := strconv.Itoa(rand.Intn(100))
		arr1[i] = val
		arr2[i] = val
	}
	if !ArrayContain(arr1, arr2) {
		t.Fatalf("array full contain with order error")
	}
	if !ArrayContain(arr1, arr2[:5]) {
		t.Fatalf("array contain with order error")
	}
	if ArrayContain(arr1, append(arr2[:5], "101")) {
		t.Fatalf("array contain with order error: shouldn't contain")
	}
}

func Test_XthSlice(t *testing.T) {
	arr1 := make([]string, 100)
	result := make(map[int]int)
	chunk := 10
	for i := 0; i < 100; i++ {
		arr1[i] = strconv.Itoa(i)
		result[i/chunk] += i
	}
	for xth, total := range result {
		arr := XthSlice(arr1, chunk, xth)
		sum := 0
		for _, ns := range arr {
			n, _ := strconv.Atoi(ns)
			sum += n
		}
		if sum != total {
			t.Fatalf("xth slice error: not match")
		}
	}
	arr := XthSlice(arr1, chunk, 11)
	if len(arr) > 0 {
		t.Fatalf("xth slice error: out of range")
	}
}
