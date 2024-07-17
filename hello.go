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

type counter struct {
	count uint32
}

func (c *counter) get() uint32 {
	c.count++
	return c.count
}
func main() {
	// Define a Student label

	// Regex "^[A-Z][a-z]+$"
	// Means Start Capital letter followed by multiple lowercase letters then End
	// "Patrick" is allowed, "Matthew" is allowed, "Helen" is allowed.
	// " Patrick" is not allowed because it starts with a space
	// "Pat rick" is not allowed because it has a space where there should be a letter
	// "Patrick1" is not allowed because it contains a digit.
	// Note: "Su Li" would NOT be allowed.

	// Regex "^(19|20)\\d{2}-(0|1)\\d-(0|1|2|3)\\d$"
	// Means start with 19 or 20 followed by 2 digits then a dash then o or 1, a digit, a dash
	// followed by 0,1,2 or 3 ending with a digit.
	// Note: "1923-13-39" would be allowed but is still an invalid date

	// Regex "^\\d{10}$"
	// Means Start 10 ditits End.

	StudentLabel := mynodes.NewLabel("Student", []mynodes.Attribute{
		{Name: "forename", Text: "First name", Required: true, Regex: "^[A-Z][a-z]+$"},
		{Name: "surname", Text: "Last name", Required: true, Regex: ""},
		{Name: "date_of_birth", Text: "Date of Birth", Required: true, Regex: "^(19|20)\\d{2}-(0|1)\\d-(0|1|2|3)\\d$"},
		{Name: "uln", Text: "ULN", Required: false, Regex: "^\\d{10}$"},
	})
	StudentLabelPtr := StudentLabel.GetPointer()

	CurrentStudentLabel := mynodes.NewLabel("Current Student", []mynodes.Attribute{})
	CurrentStudentLabelPtr := CurrentStudentLabel.GetPointer()

	var labs = []mynodes.LabelPtr{StudentLabelPtr, CurrentStudentLabelPtr}

	var props = make(map[string]string)
	props["forename"] = "Patrick"
	props["surname"] = "Biggs"
	props["date_of_birth"] = "1967-01-31"
	s1, err := mynodes.NewNode(labs, props)
	if err != nil {
		panic(err)
	}
	props["forename"] = "Matthew"
	props["date_of_birth"] = "1995-06-29"
	s2, err := mynodes.NewNode(labs, props)
	s2.Print()
	if err != nil {
		panic(err)
	}
	props["forename"] = "Helen"
	props["date_of_birth"] = "1964-07-03"
	s2, err = mynodes.NewNode(labs, props)
	s2.Print()
	if err != nil {
		panic(err)
	}
	CurrentStudent, err := mynodes.NewNode([]mynodes.LabelPtr{CurrentStudentLabelPtr}, nil)
	if err != nil {
		panic(err)
	}

	r1 := mynodes.NewRelation()
	r1.SetAB(s1, CurrentStudent)

	s1.AddProperty("uln", "1234567890")

	//mynodes.SaveAll()

	mynodes.LoadAll()
}
