package main

import "fmt"

func main() {
	aMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}

	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50
	for key, val := range m {
		fmt.Print("[ ",key," -> ", val," ]")
	}
	fmt.Println("Apple price:", m["Apple"])
	delete(m, "Apple")
	value, ok := m["Apple"]
	fmt.Println("Apple price:", value, "Present:", ok)
	value2, ok2 := m["Banana"]
	fmt.Println("Banana price:", value2, "Present:", ok2)
}


