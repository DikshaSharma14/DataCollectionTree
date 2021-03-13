package openapi

import (
	"encoding/json"
	"net/http"
)

// A InsertApiController binds http requests to an api service and writes the service results to the http response
type InsertApiController struct {
	service InsertApiServicer
}

// NewInsertApiController creates a default api controller
func NewInsertApiController(s InsertApiServicer) Router {
	return &InsertApiController{service: s}
}

// Routes returns all of the api route for the InsertApiController
func (c *InsertApiController) Routes() Routes {
	return Routes{
		{
			"Insert",
			http.MethodPut,
			"/v1/insert",
			c.Insert,
		},
	}
}

// Insert - insert a new metric
func (c *InsertApiController) Insert(w http.ResponseWriter, r *http.Request) {
	insertRequest := &InsertRequest{}
	if err := json.NewDecoder(r.Body).Decode(&insertRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.Insert(r.Context(), *insertRequest)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeResponse(&result.Code, w)

}
