package main

/*
Create a graph database capable of running a college database.
Patrick Biggs
Matthew Biggs
2024-07-09
*/

import (
	"github.com/hello/mynodes"
)

func main() {
	/* Define a Student label */
	var forename mynodes.Property = mynodes.Property{Name: "forename", Required: true, Regex: "^[A-Z][a-z]+$"}
	var surname mynodes.Property = mynodes.Property{Name: "surname", Required: true, Regex: ""}
	var uln mynodes.Property = mynodes.Property{Name: "uln", Required: false, Regex: ""}

	var Student mynodes.Label = mynodes.Label{
		LabelId:    1,
		Name:       "Student",
		Properties: []*mynodes.Property{&forename, &surname, &uln},
	}
	var CurrentStudentLabel mynodes.Label = mynodes.Label{
		LabelId: 2,
		Name:    "Current Student",
	}

	var labs = []*mynodes.Label{&Student, &CurrentStudentLabel}
	var props = []mynodes.KeyValuePair{
		{Key: "forename", Value: "Patrick"},
		{Key: "surname", Value: "Biggs"},
	}
	Steve := mynodes.NewNode(labs, props)
	Steve.Save()

	CurrentStudent := mynodes.NewNode([]*mynodes.Label{&CurrentStudentLabel}, nil)
	CurrentStudent.Save()

	r1 := mynodes.NewRelation()
	r1.SetAB(Steve, CurrentStudent)

	r2 := mynodes.NewRelation()
	r2.SetAB(Steve, CurrentStudent)

	Steve.AddProperty(mynodes.KeyValuePair{Key: "uln", Value: "0000000000"})
	Steve.Save()
}
