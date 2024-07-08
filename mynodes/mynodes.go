package mynodes

/*
This package is to implement a simple graph of nodes and relations.
Each node can have many labels.
Each node can have many relations.
Each node can have many key value pairs
Each relation should have a name
Each relation can have many key value pairs.const
*/

import "fmt"

// "forename":"Patrick"
type keyValuePair struct {
	key   string
	value string
}

type Relation struct {
	name       string          // Mandatory name
	a          *Node           // Pointer to the node the relation is from
	b          *Node           // Pointer to the node the relation is going to
	properties []*keyValuePair // List of key value pairs
}

type Node struct {
	labels     []*string      // Slice of pointers to labels
	properties []keyValuePair // I want o include a slice of key value pairs here
	relations  []*Relation    // Slice of pointers to Relations
}

func Hello() {
	fmt.Println("Hello from mynodes")
}

var labels = [4]string{"Student", "Employee", "Department", "Subject"}

func PrintLabels() {
	fmt.Println(labels)
}

//var nodeList = [1]Node

// Add a new label to the labels slice
func AddLabel(label string) {
	// Todo: check to ensure label is not already defined ebfore continuing
	var l []string = labels[0:]
	l = append(l, label)
}
