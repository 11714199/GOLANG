package main

import "fmt"

func sum(nums ...int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}

	return num
}

func changeSlice(nums ...int) {
	nums[2] = 1234
}

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	mySlice := arr[1:6]
	mySlice = mySlice[1:]
	fmt.Println(mySlice)

	//make
	mySlice1 := make([]int, 10)
	mySlice2 := make([]int, 5, 10)
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))

	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	// Variadic Functions
	total := sum(1, 2, 3, 4)
	fmt.Println(total)

	// Spread Operator
	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))

	//Append
	nums = append(nums, 5)
	fmt.Println(nums)

	nums = append(nums, 6, 7)
	fmt.Println(nums)

	mySlice = append(mySlice, nums...)
	fmt.Println(mySlice, arr)

	// Slice of Slices
	// matrix or 2D Slice
	matrix := [][]int{}

	for i := 0; i < 10; i++ {
		row := []int{}
		for j := 0; j < 10; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	fmt.Println(matrix)

	// Tricky Slice

	//with length
	tricky := make([]int, 3)
	a := append(tricky, 1)
	fmt.Println(a) // [0,0,0,1]
	b := append(tricky, 2)
	fmt.Println(b, a) // [0,0,0,2] [0,0,0,1]

	// with cap adn len
	// Cap will create new underlying array with same reference
	tricky1 := make([]int, 3, 10)
	c := append(tricky1, 1)
	fmt.Println(c) // [0,0,0,1]
	d := append(tricky1, 2)
	fmt.Println(d, c) // [0,0,0,2] [0,0,0,2]

	// range := iterate elements over a slice

	for index, ele := range mySlice {
		fmt.Println(index, ele)
	}

	// Slice is pass by reference
	changeSlice(mySlice...)
	fmt.Println(mySlice)

}
