package openapi

type QueryResponse struct {
	Dim []KeyValueString `json:"dim,omitempty"`

	Metrics []KeyValueInt `json:"metrics,omitempty"`
}
