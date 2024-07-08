package main

/*
The purpose of this program is to create a graph database capable of running a college database.

*/
import (
	"fmt"

	"github.com/hello/mynodes"
)

func hello(name string) string {
	return "Hello " + name
}

type person struct {
	forename string
	surname  string
}

func newPerson(forename string, surname string) *person {
	p := person{forename: forename, surname: surname}
	return &p
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(hello("Bozo"))

	peter := person{}
	peter.forename = "peter"
	peter.surname = "parker"
	fmt.Println(peter)
	rita := person{"Rita", "Skeeter"}
	fmt.Println(rita)
	bob := person{forename: "Bob", surname: "Silver"}
	fmt.Println(bob)
	fmt.Println(bob.forename + " " + bob.surname)

	sue := *newPerson("Sue", "McGrew")
	fmt.Println(sue)

	// Make an array of persom
	var a [4]person
	a[0] = peter
	a[1] = rita
	a[2] = bob
	a[3] = sue

	//Make a slice of the entire array
	var s []person = a[0:]

	// Make a new person record
	mark := person{"Marc", "Bolan"}

	// Add new person record to the slice
	s = append(s, mark)

	// Add a new person record to the slice directly
	s = append(s, person{"James", "Web"})

	fmt.Println(a) // Print the array
	fmt.Println(s) // Print the slice

	type people []person
	peeps := people{peter, rita, bob, sue}
	fmt.Println(peeps)

	mynodes.Hello() // Just here to prove I can access function in the mynodes package.

	mynodes.PrintLabels()  // Display the currently defined labels
	mynodes.AddLabel("PC") // Add a new label
	mynodes.PrintLabels()  // Display the labels.
	// Note this doesn't work
}
