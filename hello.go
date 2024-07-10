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
	/* Student node via raw commands */
	Steve := mynodes.NewNode()
	Steve.AddLabel("Student")
	Steve.AddLabel("Moron")
	Steve.AddProperty(mynodes.KeyValuePair{Key: "forename", Value: "Steve"})
	Steve.AddProperty(mynodes.KeyValuePair{Key: "surname", Value: "Peters"})
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
