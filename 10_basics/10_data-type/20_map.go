package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["Apple"] = 40
	map1["Banana"] = 30
	map1["Mango"] = 50
	for key, val := range map1 {
		fmt.Print("[ ",key," -> ", val," ]")
	}
	fmt.Println("Apple price:", map1["Apple"])
	delete(map1, "Apple")
	value, ok := map1["Apple"]
	fmt.Println("Apple price:", value, "Present:", ok)
	value2, ok2 := map1["Banana"]
	fmt.Println("Banana price:", value2, "Present:", ok2)

	map2 := map[string]int{
		"a": 81,
		"b": 82,
		"c": 83,
	}
	fmt.Println("map2", map2)
}


