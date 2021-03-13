package openapi

type InsertRequest struct {
	Dim []KeyValueString `json:"dim,omitempty"`

	Metrics []KeyValueInt `json:"metrics,omitempty"`
}
