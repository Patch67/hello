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
	"errors"
	"fmt"
)

var labels []Label = []Label{}          // Empty slice to hold all labels
var nodes []node = []node{}             // Empty slice to hold all Nodes
var names []string = []string{}         // Empty slice to hold all Names
var relations []relation = []relation{} // Empty slice to hold all Relations

var labelId uint32 = 1    // Unique label id
var nodeId uint32 = 1     // Unique node id
var nameId uint32 = 1     // Unique name id
var relationId uint32 = 1 // Unique relation id

type RelationPtr struct {
	RelationId  uint32
	RelationPtr *relation
}

type relation struct {
	relationId uint32
	name       string            // Mandatory name
	a          *node             // Pointer to the node the relation is from
	b          *node             // Pointer to the node the relation is going to
	properties map[string]string // List of key value pairs
}

type NodePtr struct {
	NodeId  uint32 `json:"node_id"`
	NodePtr *node  `json:"-"`
}

type node struct {
	nodeId     uint32            // Unique ID of node
	labels     []*Label          // Slice of pointers to labels
	properties map[string]string // Slice of key value pairs
	relations  []*RelationPtr    // Slice of pointers to Relations
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
	Properties []Property
}

func NewLabel(name string, properties []Property) *Label {
	newLabel := Label{Name: name, Properties: properties}
	return &newLabel
}

/* Create a new Node */
func NewNode(l []*Label, data map[string]string) *node {
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

func (node *node) AddProperty(key string, value string) {
	node.properties[key] = value
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

/* Get the value of a property */
func (node *node) Get(p string) string {
	return node.properties[p]
}

/* Set the value of a property */
func (node *node) Set(key string, value string) {
	node.properties[key] = value
}

/* Checks if the node has all the required properties */
func (node *node) IsValid() (bool, error) {
	for _, label := range node.labels {
		drl := *label // Dereferenced label cos can't range over pointer to slice
		for _, property := range drl.Properties {
			_, ok := node.properties[property.Name]
			if !ok && property.Required {
				return false, errors.New("Required property " + property.Name + " not found")
			}
		}
	}
	return true, nil
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
	count := 0
	for key, value := range node.properties {
		fmt.Print("\"", key, "\":\"", value, "\"")
		if count < len(node.properties)-1 {
			fmt.Print(",")
		}
		count++
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
