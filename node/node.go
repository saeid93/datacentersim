package node

import (
	"fmt"
)

type node struct {
	location string
	cluster  string
	capacity int
	usage    int
}

func New(location string, cluster string, capacity int, usage int) node {
	e := node{location, cluster, capacity, usage}
	return e
}

func (e node) AddedNode() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.location, e.cluster, (e.capacity - e.usage))
}
