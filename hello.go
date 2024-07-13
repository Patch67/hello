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
		{Name: "forename", Required: true, Regex: "^[A-Z][a-z]+$"},
		{Name: "surname", Required: true, Regex: ""},
		{Name: "date_of_birth", Required: true, Regex: "^(19|20)\\d{2}-(0|1)\\d-(0|1|2|3)\\d$"},
		{Name: "uln", Required: false, Regex: "^\\d{10}$"},
	})

	CurrentStudentLabel := mynodes.NewLabel("Current Student", []mynodes.Attribute{})
	var labs = []*mynodes.Label{StudentLabel, CurrentStudentLabel}

	var props = make(map[string]string)
	props["forename"] = "Patrick"
	props["surname"] = "Biggs"
	props["date_of_birth"] = "1967-11-31"

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
