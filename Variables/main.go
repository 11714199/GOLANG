package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
}

func add(x, y int) (a, b int) {

	if x == 0 {
		a = 0
	} else if x == 1 {
		a = 1
	} else {
		b = 2
		a = 1
	}

	return
}
