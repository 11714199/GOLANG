package main

import "fmt"

func changeMap(m map[string]int) {
	m["Madhavi"] = 22
}

func main() {
	var list map[string]int
	fmt.Println(list)

	maps := make(map[string]int)
	maps1 := map[string]int{
		"Madhavi": 23,
		"Madhu":   23,
	}

	fmt.Println(maps, maps1)

	// Map is pass by reference
	changeMap(maps1)
	fmt.Println(maps1)

	// Nested Maps
	// var nestedMap map[string]map[string]int
	// nestedMap["a"]["Madhavi"] = 23
	// nestedMap["b"]["Madhu"] = 22
	// for key, val := range nestedMap {
	// 	fmt.Println(key, val)
	// }

	nestedMap := make(map[string]map[string]int)
	nestedMap["c"]["Madhavi Reddy"] = 21

	for _, val := range nestedMap {
		for key1, val1 := range val {
			fmt.Printf("%v: %v", key1, val1)
			fmt.Println()
		}
	}
	fmt.Println(nestedMap)

	// nestedMap["c"]["Madhavi Reddy"] = 21
	// fmt.Println(nestedMap)
}
