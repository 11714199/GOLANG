package main

import "fmt"

type employeeDetails interface {
	getEmployeeDetails() string
	//getData() string
}

type personalDetails struct {
	first_name string
	last_name  string
}

func (p personalDetails) getEmployeeDetails() string {
	return p.first_name + " " + p.last_name
}

type address struct {
	address string
	pincode int
}

func (a address) getEmployeeDetails() string {
	return a.address
}

func getDetails(e employeeDetails) string {
	return e.getEmployeeDetails()
}

func getData(e []employeeDetails) string {
	var details string
	for _, emp := range e {
		details += emp.getEmployeeDetails() + "\n"
	}
	return details
}

func emptyInterface(value interface{}) {
	fmt.Println(value)
}

var value interface{} = 42

func assertion() {
	if val, okk := value.(int); okk {
		fmt.Println(okk)
		fmt.Println(val)
	} else {
		fmt.Println("Not a int")
	}
}

func main() {
	per := personalDetails{
		first_name: "Madhavi",
		last_name:  "Asam",
	}

	add := address{
		address: "Kadiri",
		pincode: 515581,
	}

	per1 := personalDetails{
		first_name: "Madhavi",
		last_name:  "Reddy",
	}

	//Abstraction

	emps := []employeeDetails{per, per1}
	fmt.Println(getData(emps))

	fmt.Println(getDetails(per))
	fmt.Println(getDetails(add))
	fmt.Println(getDetails(per1))

	// Empty Interface
	emptyInterface("Madhavi")
	emptyInterface(12234)
	emptyInterface(12.23)

	// Aseertion
	assertion()
}
