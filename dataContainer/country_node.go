package dataContainer

import (
	"sync"

	"github.com/DikshaSharma14/data-collection-tree/util"
)

type CountryNode struct {
	CountryName              string
	UpdateLock               sync.Mutex
	AggregatedCountryMatrics map[string]int
	Devices                  map[string]*DeviceNode
}

func NewCountryNode() *CountryNode {
	var cn CountryNode
	cn.AggregatedCountryMatrics = make(map[string]int)
	cn.Devices = make(map[string]*DeviceNode)
	return &cn
}

func (cn *CountryNode) UpdateTransaction(root *RootNode, deltaMatrics map[string]int) {
	root.UpdateLock.Lock()
	root.AggregatedWorldMetrics = util.AddMapBToA(root.AggregatedWorldMetrics, deltaMatrics)
	root.UpdateLock.Unlock()
}
