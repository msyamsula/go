package main

import "fmt"

// struct definition
// struct default is zero values, string="", integer=0, bool=false
// embedding contact to person
type person struct {
	firstName   string
	lastName    string
	contactInfo contact
}

type contact struct {
	email string
	zip   int
}

// receiver function for person
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func main() {
	alex_contact := contact{"alex@baros.com", 1234}
	alex := person{"alex", "anderson", alex_contact}
	alex.firstName = "waw"
	fmt.Println(alex)       // simple print
	fmt.Printf("%+v", alex) // print with its variable name

	fmt.Println()

	// declare with json style
	budi := person{
		firstName: "budi",
		lastName:  "enak",
		contactInfo: contact{
			email: "budi@anduk.com",
			zip:   87767,
		},
	}

	fmt.Printf("%+v\n", budi)

	// using function receiver
	budi.print()
}
