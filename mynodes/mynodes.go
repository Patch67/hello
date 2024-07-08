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
type KeyValuePair struct {
	Key   string
	Value string
}

type Relation struct {
	name       string          // Mandatory name
	a          *Node           // Pointer to the node the relation is from
	b          *Node           // Pointer to the node the relation is going to
	properties []*KeyValuePair // List of key value pairs
}

type Node struct {
	Labels     []string       // Slice of pointers to labels
	Properties []KeyValuePair // I want to include a slice of key value pairs here
	Relations  []*Relation    // Slice of pointers to Relations
}

func AddLabel(node Node, label string) {
	node.Labels = append(node.Labels, label)
	fmt.Println(node.Labels)
}

func AddProperty(node Node, property KeyValuePair) {
	node.Properties = append(node.Properties, property)
	fmt.Println(node.Properties)
}

func AddRelation(node Node, relation Relation) {
	node.Relations = append(node.Relations, &relation)
}

func PrintNode(node Node) {
	fmt.Println(node.Labels)
	fmt.Println(node.Properties)
	fmt.Println(node.Relations)
}

var labels []string = []string{} // THIS IS HOW TO DO SLICES OUTSIDE A FUNCTION!

// Create an empty slice to hold all of our Nodes
var nodes []Node = []Node{}

func PrintLabels() {
	fmt.Println(labels)
}

// Add a new label to the labels slice
func AddTestLabel(label string) {
	// Todo: check to ensure label is not already defined ebfore continuing
	labels = append(labels, label)
}
