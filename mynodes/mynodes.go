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

var labels []string = []string{}        // Empty slice to hold all labels
var nodes []Node = []Node{}             // Empty slice to hold all Nodes
var names []string = []string{}         // Empty slice to hold all Nodes
var relations []Relation = []Relation{} // Empty slice to hold all Relations

var labelId uint32 = 1
var nodeId uint32 = 1
var nameId uint32 = 1
var relationId uint32 = 1

// "forename":"Patrick"
type KeyValuePair struct {
	Key   string
	Value string
}

type RelationPtr struct {
	RelationId  uint32
	RelationPtr *Relation
}

type Relation struct {
	name       string          // Mandatory name
	a          *Node           // Pointer to the node the relation is from
	b          *Node           // Pointer to the node the relation is going to
	properties []*KeyValuePair // List of key value pairs
}

type NodePtr struct {
	NodeId  uint32
	NodePtr *Node
}

type Node struct {
	nodeId     uint32         // Unique ID of node
	labels     []string       // Slice of pointers to labels
	properties []KeyValuePair // I want to include a slice of key value pairs here
	relations  []*RelationPtr // Slice of pointers to Relations
}

/* Create a new Node */
func NewNode() *Node {
	node := Node{}
	node.nodeId = nodeId
	nodeId += 1
	nodes = append(nodes, node)
	return &node
}

/* node.AddLabel methos */
func (node *Node) AddLabel(label string) {
	node.labels = append(node.labels, label)
}

func (node *Node) AddProperty(property KeyValuePair) {
	node.properties = append(node.properties, property)
}

func (node *Node) AddRelation(relation Relation) {
	node.relations = append(node.relations, &RelationPtr{RelationId: 0, RelationPtr: &relation})
}

func (node *Node) Print() {
	fmt.Println(node.nodeId)
	fmt.Println(node.labels)
	fmt.Println(node.properties)
	fmt.Println(node.relations)
}

func (relation *Relation) SetAB(A *Node, B *Node) {
	relation.a = A
	relation.b = B
}
func (relation *Relation) Print() {
	fmt.Println(relation)
}
func PrintLabels() {
	fmt.Println(labels)
}

// Add a new label to the labels slice
func AddTestLabel(label string) {
	// Todo: check to ensure label is not already defined ebfore continuing
	labels = append(labels, label)
}
