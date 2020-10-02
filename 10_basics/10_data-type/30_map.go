package main

import "fmt"

func main() {
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


