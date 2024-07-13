package main

/*
Create a graph database capable of running a college database.
Patrick Biggs
Matthew Biggs
2024-07-09
*/

import (
	"fmt"

	"github.com/hello/mynodes"
)

func main() {
	// Define a Student label
	StudentLabel := mynodes.NewLabel("Student", []mynodes.Attribute{
		{Name: "forename", Required: true, Regex: "^[A-Z][a-z]+$"},
		{Name: "surname", Required: true, Regex: ""},
		{Name: "uln", Required: false, Regex: "^\\d{10}$"},
	})

	CurrentStudentLabel := mynodes.NewLabel("Current Student", []mynodes.Attribute{})
	var labs = []*mynodes.Label{StudentLabel, CurrentStudentLabel}

	var props = make(map[string]string)
	props["forename"] = "Patrick"
	props["surname"] = "Biggs"

	Steve, err := mynodes.NewNode(labs, props)
	if err != nil {
		panic(err)
	}
	Steve.Save()

	CurrentStudent, err := mynodes.NewNode([]*mynodes.Label{CurrentStudentLabel}, nil)
	if err != nil {
		panic(err)
	}
	CurrentStudent.Save()
	fmt.Println(CurrentStudent.IsValid())

	r1 := mynodes.NewRelation()
	r1.SetAB(Steve, CurrentStudent)

	r2 := mynodes.NewRelation()
	r2.SetAB(Steve, CurrentStudent)

	Steve.AddProperty("uln", "1234567890")
	Steve.Save()
	fmt.Println(Steve.Get("forename"))
	fmt.Println(Steve.Get("bozo"))
	fmt.Println(Steve.IsValid())

	// Let's try creating a student without all the required properties
	bob, err := mynodes.NewNode([]*mynodes.Label{StudentLabel, CurrentStudentLabel}, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bob)
}
