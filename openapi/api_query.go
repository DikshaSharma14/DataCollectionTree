package openapi

import (
	"encoding/json"
	"net/http"
)

// A QueryApiController binds http requests to an api service and writes the service results to the http response
type QueryApiController struct {
	service QueryApiServicer
}

// NewQueryApiController creates a default api controller
func NewQueryApiController(s QueryApiServicer) Router {
	return &QueryApiController{service: s}
}

// Routes returns all of the api route for the QueryApiController
func (c *QueryApiController) Routes() Routes {
	return Routes{
		{
			"Query",
			http.MethodPost,
			"/v1/query",
			c.Query,
		},
	}
}

// Query - queries the mertic
func (c *QueryApiController) Query(w http.ResponseWriter, r *http.Request) {
	queryRequest := &QueryRequest{}
	if err := json.NewDecoder(r.Body).Decode(&queryRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.Query(r.Context(), *queryRequest)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
