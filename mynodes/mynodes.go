package mynodes

/*
This package implements a simple graph of nodes and relations.
Each node can have many labels.
Each node can have many relations.
Each node can have many key value pairs

Each label can have many attributes

Each relation should have a name
Each relation can have many key value pairs

*/

import (
	"errors"
	"fmt"
	"regexp"
)

var labels []Label = []Label{}          // Empty slice to hold all labels
var nodes []node = []node{}             // Empty slice to hold all Nodes
var names []string = []string{}         // Empty slice to hold all Names
var relations []relation = []relation{} // Empty slice to hold all Relations

var nameId uint32 = 1 // Unique name id

// closure stores a var called 'next' which gets incremented and returned when called
var NewLabelId = func() func() uint32 {
	var next uint32 = 0
	return func() uint32 {
		next++
		return next
	}
}()

var nextNodeID = func() func() uint32 {
	var next uint32 = 0
	return func() uint32 {
		next++
		return next
	}
}()

var nextRelationID = func() func() uint32 {
	var next uint32 = 0
	return func() uint32 {
		next++
		return next
	}
}()

type RelationPtr struct {
	RelationId  uint32    `json:"relation_id,omitempty"`
	RelationPtr *relation `json:"relation_ptr,omitempty"`
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
type Attribute struct {
	Name     string
	Required bool
	Regex    string
}

type Label struct {
	LabelId    uint32
	Name       string
	Attributes []Attribute
}

func NewLabel(name string, attributes []Attribute) *Label {
	newLabel := Label{LabelId: NewLabelId(), Name: name, Attributes: attributes}
	labels = append(labels, newLabel)
	return &newLabel
}

/* Create a new Node */
func NewNode(l []*Label, data map[string]string) (*node, error) {
	//Create node
	newNode := node{nodeId: nextNodeID(), labels: l, properties: data}
	v, e := newNode.IsValid()
	if !v {
		return nil, errors.New(e.Error())
	}
	nodes = append(nodes, newNode)
	return &newNode, nil
}

/* Add label to node */
func (node *node) AddLabel(label *Label) {
	node.labels = append(node.labels, label)
}

/* Add proerty to node */
func (node *node) AddProperty(key string, value string) {
	node.properties[key] = value
}

/* Add relation to node */
func (node *node) AddRelation(relation relation) {
	node.relations = append(node.relations, &RelationPtr{RelationId: 0, RelationPtr: &relation})
}

/* Display node */
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
		for _, attribute := range label.Attributes {
			_, ok := node.properties[attribute.Name] // Does the property exist
			if !ok && attribute.Required {           // Is the property required?
				return false, errors.New("Required property " + attribute.Name + " not found")
			}
			// If there is regex in the attribute does the value conform
			if ok && attribute.Regex != "" {
				r, _ := regexp.Compile(attribute.Regex)
				if !r.MatchString(node.properties[attribute.Name]) {
					return false, errors.New(attribute.Name + " value " + node.properties[attribute.Name] + " is not " + attribute.Regex)
				}
			}

		}
	}
	return true, nil
}

/* Save a node to screen */
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

// Create a new relation
func NewRelation() *relation {
	relation := relation{relationId: nextRelationID()}
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
