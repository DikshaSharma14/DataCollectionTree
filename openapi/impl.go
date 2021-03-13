package openapi

//Implementation response defines an error code with the associated body
type ImplResponse struct {
	Code int
	Body interface{} `json:"omitempty"`
}
