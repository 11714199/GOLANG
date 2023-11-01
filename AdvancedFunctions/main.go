package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func getSum(a, b int, add func(int, int) int) int {
	return add(a, b)
}

func curriedFun(x, z int) func(int) int {
	return func(y int) int {
		return x + y + z
	}
}

func closures() func() int {
	count := 0 // This variable is captured by the closure
	return func() int {
		count++
		return count
	}
}
func main() {
	// first-order functions
	sum := add(1, 2)
	fmt.Println(sum)

	// High-order functions
	getSum := getSum(2, 3, add)
	fmt.Println(getSum)

	// function curring
	getX := curriedFun(2, 3)
	getY := getX(3)
	fmt.Println(getY)

	// defer
	defer fmt.Println("With Defer")
	fmt.Println("Withour defer")

	//Closures
	getClosure := closures()
	fmt.Println(getClosure()) // count 1
	fmt.Println(getClosure()) // count 2

	// Anonymouse Functions
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(4, 5))

	multiply := func(x, y int) int {
		return x * y
	}(3, 4)
	fmt.Println(multiply)
}
