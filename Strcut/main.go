package main

import "fmt"

// nested structure
type personal struct {
	first_name    string
	last_name     string
	mobile_number int
	address       Address
}

type Address struct {
	address  string
	landMark string
	pincode  int
}

// embedded struct

type personalDetails struct {
	first_name    string
	last_name     string
	mobile_number int
	Address
}

func main() {
	// anonymous strcut
	details := struct {
		first_name string
		last_name  string
	}{
		first_name: "Madhavi",
		last_name:  "Reddy",
	}
	fmt.Println(details)
	var personalInfo personal
	personalInfo.first_name = "Madhavi"
	personalInfo.last_name = "Asam"
	personalInfo.mobile_number = 1234567899
	personalInfo.address.address = "Kadiri"
	personalInfo.address.landMark = "Home"
	personalInfo.address.pincode = 515581

	var personalInfo1 personalDetails
	personalInfo1.first_name = "Madhavi"
	personalInfo1.last_name = "Reddy"
	personalInfo1.mobile_number = 1234567899
	personalInfo1.address = "Anantapur"
	personalInfo1.landMark = "Kadiri"
	personalInfo1.pincode = 515599

	personalInfo1.getpersonalDetails()
	personalInfo.getDetails()
	getStruct(personalInfo.first_name)
}

// receiver function
func (p personalDetails) getpersonalDetails() {
	fmt.Println(p)
}

func (p personal) getDetails() {
	fmt.Println(p)
}

func getStruct(s string) {

}
