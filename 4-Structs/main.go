package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// // give access to the original struct's address
	// jimPointer := &jim
	// jimPointer.updateFirstName("Jimmy")
	jim.updateFirstName("Jimmy")

	jim.print()
}

//* is pointer to person
func (p *person) updateFirstName(newFirstName string) {
	// // this is the operator to manipulate the value the pointer points to
	// (*p).firstName = newFirstName

	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
