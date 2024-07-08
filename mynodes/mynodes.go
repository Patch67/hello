package mynodes

import "fmt"

type keyValuePair struct {
	key   string
	value string
}

type Node struct {
	labels     []*string
	properties []keyValuePair
	relations  []*Node
}

func Hello() {
	fmt.Println("Hello from mynodes")
}

//nodeList := []*Node
