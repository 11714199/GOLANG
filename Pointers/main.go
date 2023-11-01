package main

import "fmt"

type employee struct {
	first_name string
	last_name  string
}

// Ponter receiver function
func (e *employee) changesNameWithPointer() {
	e.last_name = "Asam"
}

// non Ponter receiver function
func (e employee) changesName() {
	e.last_name = "Asam"
}
func main() {
	var p *int
	var x int
	fmt.Println(x, p)
	p = &x
	fmt.Println(x, *p)

	x = 10
	fmt.Println(x, *p)
	*p = 1234
	fmt.Println(x, *p)

	// Ponter receiver
	emp := employee{
		first_name: "Madhavi",
		last_name:  "Reddy",
	}
	emp.changesName()
	fmt.Println(emp) //{Madhavi Reddy}

	emp.changesNameWithPointer()
	fmt.Println(emp) //{Madhavi Asam}
}
