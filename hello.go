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
	/* Define a Student label */
	var forename mynodes.Attribute = mynodes.Attribute{Name: "forename", Required: true, Regex: "^[A-Z][a-z]+$"}
	var surname mynodes.Attribute = mynodes.Attribute{Name: "surname", Required: true, Regex: ""}
	var uln mynodes.Attribute = mynodes.Attribute{Name: "uln", Required: false, Regex: ""}

	var Student mynodes.Label = mynodes.Label{
		LabelId:    1,
		Name:       "Student",
		Properties: []mynodes.Attribute{forename, surname, uln},
	}
	var CurrentStudentLabel mynodes.Label = mynodes.Label{
		LabelId: 2,
		Name:    "Current Student",
	}

	var labs = []*mynodes.Label{&Student, &CurrentStudentLabel}

	var props = make(map[string]string)
	props["forename"] = "Patrick"
	props["surname"] = "Biggs"

	Steve, err := mynodes.NewNode(labs, props)
	if err != nil {
		panic(err)
	}
	Steve.Save()

	CurrentStudent, err := mynodes.NewNode([]*mynodes.Label{&CurrentStudentLabel}, nil)
	if err != nil {
		panic(err)
	}
	CurrentStudent.Save()
	fmt.Println(CurrentStudent.IsValid())

	r1 := mynodes.NewRelation()
	r1.SetAB(Steve, CurrentStudent)

	r2 := mynodes.NewRelation()
	r2.SetAB(Steve, CurrentStudent)

	Steve.AddProperty("uln", "0000000000")
	Steve.Save()
	fmt.Println(Steve.Get("forename"))
	fmt.Println(Steve.Get("bozo"))
	fmt.Println(Steve.IsValid())

	// Let's try creating a student without all the required properties
	bob, err := mynodes.NewNode([]*mynodes.Label{&Student, &CurrentStudentLabel}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(bob.IsValid()) // Is it valid?
}
