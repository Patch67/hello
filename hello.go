package main

/*
The purpose of this program is to create a graph database capable of running a college database.

*/
import (
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
	/*
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

		fmt.Println("Array - ", a) // Print the array
		fmt.Println("Slice - ", s) // Print the slice

		peeps := []person{} // THIS IS HOW TO DO SLICES IN A FUNCTION
		peeps = append(peeps, peter)
		peeps = append(peeps, rita)
		peeps = append(peeps, bob)
		peeps = append(peeps, sue)
		peeps = append(peeps, mark)
		fmt.Println("Working -", peeps)

		mynodes.PrintLabels()           // Display the currently defined labels
		mynodes.AddTestLabel("Student") // Add a new label
		mynodes.AddTestLabel("Employee")
		mynodes.PrintLabels() // Display the labels.
	*/
	/*
		John := mynodes.Node{
			Labels: []string{"Student"},
			Properties: []mynodes.KeyValuePair{
				{
					Key:   "forename",
					Value: "string",
				},
				{
					Key:   "surname",
					Value: "string",
				},
			},
			Relations: []*mynodes.RelationPtr{},
		}
		fmt.Println(John)
		mynodes.AddLabel(John, "Student")
		mynodes.AddProperty(John, mynodes.KeyValuePair{Key: "forename", Value: "John"})
		mynodes.AddProperty(John, mynodes.KeyValuePair{Key: "surname", Value: "Smith"})
		fmt.Println(John)
		mynodes.PrintNode(John)
	*/
	Steve := mynodes.Node{}
	Steve.AddLabel("Student")
	Steve.AddProperty(mynodes.KeyValuePair{Key: "forename", Value: "Steve"})
	Steve.AddProperty(mynodes.KeyValuePair{Key: "surname", Value: "Peters"})
	Steve.Print()

	CurrentStudent := mynodes.Node{}
	CurrentStudent.AddLabel("Current Student")
	CurrentStudent.Print()

	IsCurrent := mynodes.Relation{}
	IsCurrent.SetAB(&Steve, &CurrentStudent)
	IsCurrent.Print()
}
