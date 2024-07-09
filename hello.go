package main

/*
The purpose of this program is to create a graph database capable of running a college database.

*/
import (
	"github.com/hello/mynodes"
)

func newStudent(forename string, surname string) *mynodes.Node {
	student := mynodes.NewNode()
	student.AddLabel("Student")
	student.AddProperty(mynodes.KeyValuePair{Key: "forename", Value: forename})
	student.AddProperty(mynodes.KeyValuePair{Key: "surname", Value: surname})
	return student
}

func main() {
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

	Bart := newStudent("Bart", "Simpson")
	Bart.Print()
}
