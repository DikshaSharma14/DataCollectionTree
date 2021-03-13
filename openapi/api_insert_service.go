package openapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DikshaSharma14/data-collection-tree/dataContainer"
)

var DCT *dataContainer.DataCollectionTree

func init() {
	DCT = dataContainer.CreateDataCollectionTree()
}

// InsertApiService is a service that implents the logic for the InsertApiServicer
// This service should implement the business logic for every endpoint for the InsertApi API.
// Include any external packages or services that will be required by this service.
type InsertApiService struct {
}

// NewInsertApiService creates a default api service
func NewInsertApiService() InsertApiServicer {
	return &InsertApiService{}
}

// Insert - insert a new metric
func (s *InsertApiService) Insert(ctx context.Context, insertRequest InsertRequest) (ImplResponse, error) {
	dn, country, err := getDeviceNode(insertRequest.Dim, insertRequest.Metrics)
	if err == nil {
		if _, present := DCT.DctRoot.Countries[country]; !present {
			cn := dataContainer.NewCountryNode()
			cn.CountryName = country
			DCT.InsertCountry(cn)
		}
		DCT.InsertDevice(dn, DCT.DctRoot.Countries[country])

		//dctJSON, err := json.MarshalIndent(*DCT, "", "  ")
		//if err != nil {
		//	fmt.Println(fmt.Errorf(err.Error()))
		//}
		//fmt.Printf("Marshal funnction output %s\n", string(dctJSON))

		return Response(http.StatusOK, nil), nil
	} else {
		return Response(http.StatusBadRequest, nil), err
	}
}

func getDeviceNode(dim []KeyValueString, mat []KeyValueInt) (*dataContainer.DeviceNode, string, error) {
	deviceNode := dataContainer.NewDeviceNode()
	m := make(map[string]int)
	var country string
	for i := 0; i < len(dim); i++ {
		kv := dim[i]
		if kv.Key == "country" {
			country = kv.Val
		} else if kv.Key == "device" {
			deviceNode.DeviceName = kv.Val
		}
	}
	for i := 0; i < len(mat); i++ {
		kv := mat[i]
		if kv.Val < 0 {
			return nil, "", fmt.Errorf("invalid Request Sent")
		}
		m[kv.Key] = kv.Val
	}
	deviceNode.DeviceMetrics = m
	return deviceNode, country, nil
}
