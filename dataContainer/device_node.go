package dataContainer

import (
	"sync"

	"github.com/DikshaSharma14/data-collection-tree/util"
)

type DeviceNode struct {
	DeviceName    string
	UpdateLock    sync.Mutex
	DeviceMetrics map[string]int
}

func NewDeviceNode() *DeviceNode {
	var d DeviceNode
	d.DeviceMetrics = make(map[string]int)
	return &d
}

func (dn *DeviceNode) UpdateTransaction(countryNode *CountryNode, root *RootNode, deltaMatrics map[string]int) {
	countryNode.UpdateLock.Lock()
	countryNode.AggregatedCountryMatrics = util.AddMapBToA(countryNode.AggregatedCountryMatrics, deltaMatrics)
	countryNode.UpdateTransaction(root, deltaMatrics)
	countryNode.UpdateLock.Unlock()
}
