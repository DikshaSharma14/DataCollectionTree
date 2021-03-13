package dataContainer

import (
	"sync"
)

type RootNode struct {
	UpdateLock             sync.Mutex
	AggregatedWorldMetrics map[string]int
	Countries              map[string]*CountryNode
}

func NewRootNode() *RootNode {
	var r RootNode
	r.AggregatedWorldMetrics = make(map[string]int)
	r.Countries = make(map[string]*CountryNode)
	return &r
}
