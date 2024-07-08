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

	var a [4]person
	a[0] = peter
	a[1] = rita
	a[2] = bob
	a[3] = sue

	var s []person = a[0:]

	mark := person{"Marc", "Bolan"}
	s = append(s, mark)
	s = append(s, person{"James", "Web"})
	fmt.Println(a)
	fmt.Println(s)

	type people []person
	peeps := people{peter, rita, bob, sue}
	fmt.Println(peeps)

	mynodes.Hello()

	mynodes.PrintLabels()
	mynodes.AddLabel("PC")
	mynodes.PrintLabels()
}
