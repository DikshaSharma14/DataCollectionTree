package dataContainer

import "fmt"

type DataCollectionTree struct {
	DctRoot *RootNode
}

type Operations interface {
	CreateDataCollectionTree()
	InsertCountry()
	InsertDevice()
	UpdateTransaction()
	GetAggregatedMatrics()
}

func CreateDataCollectionTree() *DataCollectionTree {
	root := NewRootNode()
	dct := &DataCollectionTree{DctRoot: root}
	return dct
}

func (dct *DataCollectionTree) InsertCountry(country *CountryNode) {
	dct.DctRoot.Countries[country.CountryName] = country
}

func (dct *DataCollectionTree) InsertDevice(device *DeviceNode, country *CountryNode) {
	deltaMatrics := make(map[string]int)
	if value, present := country.Devices[device.DeviceName]; present {
		for k, v := range device.DeviceMetrics {
			deltaMatrics[k] = v - value.DeviceMetrics[k]
		}
		device.UpdateTransaction(country, dct.DctRoot, deltaMatrics)
		country.Devices[device.DeviceName] = device
	} else {
		country.Devices[device.DeviceName] = device
		device.UpdateTransaction(country, dct.DctRoot, device.DeviceMetrics)
	}
}

func (dct *DataCollectionTree) GetAggregatedMatrics(countryName string, deviceName string) (map[string]int, error) {
	if countryNode, present := dct.DctRoot.Countries[countryName]; present {
		if deviceName != "" {
			if DeviceNode, present := dct.DctRoot.Countries[countryName].Devices[deviceName]; present {
				return DeviceNode.DeviceMetrics, nil
			} else {
				return nil, fmt.Errorf("device %s Not Found", deviceName)
			}
		}
		return countryNode.AggregatedCountryMatrics, nil
	} else {
		return nil, fmt.Errorf("country %s Not Found", countryName)
	}
}
