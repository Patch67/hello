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
import (
	"fmt"
)

var labels []string = []string{}        // Empty slice to hold all labels
var nodes []node = []node{}             // Empty slice to hold all Nodes
var names []string = []string{}         // Empty slice to hold all Nodes
var relations []relation = []relation{} // Empty slice to hold all Relations

var labelId uint32 = 1    // Unique label id
var nodeId uint32 = 1     // Unique node id
var nameId uint32 = 1     // Unique name id
var relationId uint32 = 1 // Unique relation id

// "forename":"Patrick"
type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (kv *KeyValuePair) save() {
	fmt.Print("\"", kv.Key, "\":")
	fmt.Print("\"", kv.Value, "\"")
}

type RelationPtr struct {
	RelationId  uint32
	RelationPtr *relation
}

type relation struct {
	relationId uint32
	name       string          // Mandatory name
	a          *node           // Pointer to the node the relation is from
	b          *node           // Pointer to the node the relation is going to
	properties []*KeyValuePair // List of key value pairs
}

type NodePtr struct {
	NodeId  uint32 `json:"node_id"`
	NodePtr *node  `json:"-"`
}

type node struct {
	nodeId     uint32         `json:"node_id"`    // Unique ID of node
	labels     []*Label       `json:"labels"`     // Slice of pointers to labels
	properties []KeyValuePair `json:"properties"` // Slice of key value pairs
	relations  []*RelationPtr `json:"relations"`  // Slice of pointers to Relations
}

// name: ULN,required: true,regex:"^\d{10}$"
type Property struct {
	Name     string
	Required bool
	Regex    string
}

type Label struct {
	LabelId    uint32
	Name       string
	Properties []*Property
}

func NewLabel(name string, properties []*Property) *Label {
	newLabel := Label{Name: name, Properties: properties}
	return &newLabel
}

/* Create a new Node */
func NewNode(l []*Label, data []KeyValuePair) *node {
	//Create node
	node := node{}
	node.labels = l
	node.properties = data
	node.nodeId = nodeId
	nodeId += 1
	nodes = append(nodes, node)
	return &node
}

/* node.AddLabel method */
func (node *node) AddLabel(label *Label) {
	node.labels = append(node.labels, label)
}

func (node *node) AddProperty(property KeyValuePair) {
	node.properties = append(node.properties, property)
}

func (node *node) AddRelation(relation relation) {
	node.relations = append(node.relations, &RelationPtr{RelationId: 0, RelationPtr: &relation})
}

func (node *node) Print() {
	fmt.Println(node.nodeId)
	fmt.Println(node.labels)
	fmt.Println(node.properties)
	fmt.Println(node.relations)
}

func (node *node) Save() {
	fmt.Print("{")
	fmt.Printf("\"node_id\":%d,", node.nodeId)
	fmt.Print("\"labels\":[")
	for i, value := range node.labels {
		fmt.Print("\"", value.Name, "\"")
		if i < len(node.labels)-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("],")
	fmt.Print("\"properties\":[")
	for i, value := range node.properties {
		value.save()
		if i < len(node.properties)-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("]")
	fmt.Println("}")
}

func NewRelation() *relation {
	relation := relation{}
	relation.relationId = relationId
	relationId += 1
	relations = append(relations, relation)
	return &relation
}
func (relation *relation) SetAB(A *node, B *node) {
	relation.a = A
	relation.b = B
}
func (relation *relation) Print() {
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

/* Helper functions */
/*func NewStudent(forename string, surname string) *node {
	label := Label{}
	student := NewNode(&label, "'forename':'Patrick','surname':'Biggs'")
	student.AddLabel("Student")
	student.AddProperty(KeyValuePair{Key: "forename", Value: forename})
	student.AddProperty(KeyValuePair{Key: "surname", Value: surname})
	student.AddProperty(KeyValuePair{Key: "uln", Value: "0000000000"})
	return student
}*/
