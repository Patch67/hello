package mynodes

/*
This package implements a simple graph of nodes and relations.
Each node can have many labels.
Each node can have many relations.
Each node can have many key value pairs
Each relation should have a name
Each relation can have many key value pairs

We have a nodes var which stores all nodes in a memory in a slice.
We have a relations var which stores all relations in memory in a slice.
*/

/*
Next step is to setup a label as a template so when a node is made it has the required properties.
Also set up name as a template so new relations have the required properties.
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

var labels []string = []string{}        // Empty slice to hold all labels
var nodes []Node = []Node{}             // Empty slice to hold all Nodes
var names []string = []string{}         // Empty slice to hold all Nodes
var relations []Relation = []Relation{} // Empty slice to hold all Relations

func PrintLabels() {
	fmt.Println(labels)
}

// Add a new label to the labels slice
func AddTestLabel(label string) {
	// Todo: check to ensure label is not already defined ebfore continuing
	labels = append(labels, label)
}
