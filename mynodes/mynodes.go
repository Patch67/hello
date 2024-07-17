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
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type database = struct {
	labels    []Label
	nodes     []node
	names     []string
	relations []relation
}

var db database // The God variable that holds everything in RAM

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
	NodeId     uint32            `json:"node_id,omitempty"`
	LabelPtrs  []LabelPtr        `json:"labels,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Relations  []*RelationPtr    `json:"relations,omitempty"`
}
type Node2 struct {
	NodeId uint32   `json:"node_id,omitempty"`
	Labels []*Label `json:"labels,omitempty"`
}

// name: ULN,required: true,regex:"^\d{10}$"
type Attribute struct {
	Name     string // Name of the field
	Text     string // Name of the field for UI view
	Required bool   // Is this mandatory
	Regex    string // Optionally a simple field validation string
}

/* A Label really defines a schema or a type of a node */
type Label struct {
	LabelId    uint32
	Name       string
	Attributes []Attribute
}
type LabelPtr struct {
	LabelId  uint32 `json:"label_id,omitempty"`
	LabelPtr *Label `json:"label_ptr,omitempty"`
}

func (lbl *Label) GetPointer() LabelPtr {
	newPtr := LabelPtr{lbl.LabelId, lbl}
	return newPtr
}

/* Create a new label */
func NewLabel(name string, attributes []Attribute) *Label {
	newLabel := Label{LabelId: NewLabelId(), Name: name, Attributes: attributes}
	labels = append(labels, newLabel)
	return &newLabel
}

/* Create a new Node */
func NewNode(l []LabelPtr, data map[string]string) (*node, error) {
	//Create node
	newNode := node{NodeId: nextNodeID(), LabelPtrs: l, Properties: data}
	valid, e := newNode.IsValid()
	if !valid {
		return nil, errors.New(e.Error())
	}
	nodes = append(nodes, newNode)
	return &newNode, nil
}

func ReadNode(filename string) {
	file, err := os.Open("Nodes\\" + filename)
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	newNode := make(map[string]any)
	fmt.Println(newNode)
	if err := json.Unmarshal([]byte(scanner.Text()), &newNode); err != nil {
		panic(err)
	}
	fmt.Println(newNode)
}

/* Add label to node */
func (node *node) AddLabel(label LabelPtr) {
	node.LabelPtrs = append(node.LabelPtrs, label)
}

/* Add proerty to node */
func (node *node) AddProperty(key string, value string) {
	node.Properties[key] = value
}

/* Add relation to node */
func (node *node) AddRelation(relation relation) {
	node.Relations = append(node.Relations, &RelationPtr{RelationId: 0, RelationPtr: &relation})
}

/* Display node */
func (node *node) Print() {
	fmt.Println(node.NodeId)
	fmt.Println(node.LabelPtrs)
	fmt.Println(node.Properties)
	fmt.Println(node.Relations)
}

/* Get the value of a property */
func (node *node) Get(p string) string {
	return node.Properties[p]
}

/* Set the value of a property */
func (node *node) Set(key string, value string) {
	node.Properties[key] = value
}

/* Checks if the node has all the required properties */
func (node *node) IsValid() (bool, error) {
	for _, label := range node.LabelPtrs {
		attributes := label.LabelPtr.Attributes
		for _, attribute := range attributes {
			_, ok := node.Properties[attribute.Name] // Does the property exist
			if !ok && attribute.Required {           // Is the property required?
				return false, errors.New("Required property " + attribute.Name + " not found")
			}
			// If there is regex in the attribute does the value conform
			if ok && attribute.Regex != "" {
				r, _ := regexp.Compile(attribute.Regex)
				if !r.MatchString(node.Properties[attribute.Name]) {
					return false, errors.New(attribute.Name + " value " + node.Properties[attribute.Name] + " is not " + attribute.Regex)
				}
			}

		}
	}
	return true, nil
}
func (nnode *node) Save2() string {
	/*
		Marshaling and Unmarshaling structs with pointers in them is tricky.
		Marshal deferences pointers and makes JSON with all the objects pointed to.
		This means instances that are pointed to by multiple objects are saved multiple times.
		This is inefficient.
		Many students will point to the StudentLbl but I only want to save the StudentLbl once,
		not for every single student. There could be thousands of students all pointing to
		the same StudentLbl.
		My solution is to have a struct LabelPtr that contains the pointer and the LabelId.
		When I want to convert it to JSON I will just null the pointer but keep the LabelId.
		This does mean that when reading in the JSON and contructing real objects I will need
		to search for the relevant Label by LablId then make a pointer to the found Label.

	*/
	//fmt.Println(node)

	m := make(map[string]any)
	m["node_id"] = nnode.NodeId
	// Create an array of LabelIds and JSONify that, ignoring the pointers
	lbl := []uint32{}
	for _, labelPtr := range nnode.LabelPtrs {
		lbl = append(lbl, labelPtr.LabelId)
	}
	m["labels"] = lbl
	m["properties"] = nnode.Properties
	//fmt.Println("m = ", m)
	s, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(s)
}

/* Save a node to screen */
func (node *node) Save() {
	// open output file
	file, err := os.Create("Nodes\\" + strconv.FormatUint(uint64(node.NodeId), 10) + ".txt")
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	// Print JSON to file
	fmt.Fprint(file, "{")
	fmt.Fprintf(file, "\"node_id\":%d,", node.NodeId)
	fmt.Fprint(file, "\"labels\":[")
	for i, value := range node.LabelPtrs {
		fmt.Fprint(file, "\"", value.LabelId, "\"")
		if i < len(node.LabelPtrs)-1 {
			fmt.Fprint(file, ",")
		}
	}
	fmt.Fprint(file, "],")
	fmt.Fprint(file, "\"properties\":[")
	count := 0
	for key, value := range node.Properties {
		fmt.Fprint(file, "{\"", key, "\":\"", value, "\"}")
		if count < len(node.Properties)-1 {
			fmt.Fprint(file, ",")
		}
		count++
	}
	fmt.Fprint(file, "]")
	fmt.Fprintln(file, "}")
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

func SaveAll() {
	// Save all labels
	// Save all nodes
	// Save all names
	// Save all relations
	// Save state
	file, err := os.Create("Nodes\\data.json")
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	fmt.Fprint(file, "[")
	for i, n := range nodes {
		fmt.Fprint(file, n.Save2())
		if i < len(nodes)-1 {
			fmt.Fprintln(file, ",")
		}
	}
	fmt.Fprintln(file, "]")
}

func LoadAll() {
	fmt.Println("LoadAll")
	file, err := os.Open("Nodes\\data.json")
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	var s []byte
	fmt.Fscanln(file, &s)
	s = s[1 : len(s)-1] // Missout the first [ and the last ,
	var m = make(map[string]string)
	if err := json.Unmarshal(s, &m); err != nil {
		panic(err)
	}
	fmt.Println(m)
	fmt.Println(m["labels"])

	labs := []string{}

	if err := json.Unmarshal(m["labels"].([]byte), &labs); err != nil {
		panic(err)
	}
	if rec, ok := m["labels"].(string); ok {
		for key, val := range rec {
			fmt.Printf(" [========>] %s = %v", key, val)
		}
	} else {
		fmt.Printf("record not a map[string]string: %v\n", m["labels"])
	}
	/*scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	newDb := make(map[string]any)
	fmt.Println(newDb)
	if err := json.Unmarshal([]byte(scanner.Text()), &newDb); err != nil {
		panic(err)
	}
	fmt.Println(newDb)*/
}
