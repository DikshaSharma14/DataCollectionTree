package openapi

import (
	"context"
	"fmt"
	"net/http"
)

// QueryApiService is a service that implents the logic for the QueryApiServicer
// This service should implement the business logic for every endpoint for the QueryApi API.
// Include any external packages or services that will be required by this service.
type QueryApiService struct {
}

// NewQueryApiService creates a default api service
func NewQueryApiService() QueryApiServicer {
	return &QueryApiService{}
}

// Query - queries the mertic
func (s *QueryApiService) Query(ctx context.Context, queryRequest QueryRequest) (ImplResponse, error) {
	kv := queryRequest.Dim
	var countryName string
	var deviceName string
	for i := 0; i < len(kv); i++ {
		if kv[i].Key == "country" {
			countryName = kv[i].Val
		} else if kv[i].Key == "device" {
			deviceName = kv[i].Val
		}
	}
	m, err := DCT.GetAggregatedMatrics(countryName, deviceName)
	var res *QueryResponse
	if err != nil {
		return Response(http.StatusNotFound, nil), err
	} else {
		res = createResponse(queryRequest, m)
	}
	fmt.Println(*res)
	return Response(http.StatusOK, &res), nil
}

func createResponse(qr QueryRequest, mat map[string]int) *QueryResponse {
	m := make([]KeyValueInt, 0)
	for k, v := range mat {
		m = append(m, KeyValueInt{Key: k, Val: v})
	}
	res := &QueryResponse{Dim: qr.Dim, Metrics: m}
	return res
}
