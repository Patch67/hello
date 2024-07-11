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
	var forename mynodes.Property = mynodes.Property{
		Name:     "forename",
		Required: true,
		Regex:    "^[A-Z][a-z]+$",
	}
	var Student mynodes.Label = mynodes.Label{
		LabelId:    1,
		Name:       "Student",
		Properties: []*mynodes.Property{&forename},
	}

	/* Student node via raw commands */
	Steve := mynodes.NewNode(&Student, "{'forename':'Patrick','surname':'Biggs'}")
	// The problem with adding a student this way is that we missed off ULN
	// ULN should be mandatory
	Steve.Save()

	CurrentStudent := mynodes.NewNode()
	CurrentStudent.AddLabel("Current Student")
	//CurrentStudent.Print()

	r1 := mynodes.NewRelation()
	r1.SetAB(Steve, CurrentStudent)
	//r1.Print()

	/* New Student via helper function */
	Bart := mynodes.NewStudent("Bart", "Simpson")
	//Bart.Print()
	Bart.Save()

	r2 := mynodes.NewRelation()
	r2.SetAB(Bart, CurrentStudent)
	//r2.Print()

	Bart.Save()
}
