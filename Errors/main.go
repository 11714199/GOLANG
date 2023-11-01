package main

import "fmt"

func add(a, b int) (string, error) {
	return fmt.Sprintf("Addintion of a and b is: %v", a+b), nil
}

func main() {
	res, err := add(1, 2)
	if err == nil {
		panic("Error")
		//return
	}
	fmt.Println(res)
}
