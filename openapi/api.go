package openapi

import (
	"context"
	"net/http"
)

// InsertApiRouter defines the required methods for binding the api requests to a responses for the InsertApi
// The InsertApiRouter implementation should parse necessary information from the http request,
// pass the data to a InsertApiServicer to perform the required actions, then write the service results to the http response.
type InsertApiRouter interface {
	Insert(http.ResponseWriter, *http.Request)
}

// QueryApiRouter defines the required methods for binding the api requests to a responses for the QueryApi
// The QueryApiRouter implementation should parse necessary information from the http request,
// pass the data to a QueryApiServicer to perform the required actions, then write the service results to the http response.
type QueryApiRouter interface {
	Query(http.ResponseWriter, *http.Request)
}

// InsertApiServicer defines the api actions for the InsertApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type InsertApiServicer interface {
	Insert(context.Context, InsertRequest) (ImplResponse, error)
}

// QueryApiServicer defines the api actions for the QueryApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type QueryApiServicer interface {
	Query(context.Context, QueryRequest) (ImplResponse, error)
}
