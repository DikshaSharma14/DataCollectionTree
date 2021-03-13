package openapi

type KeyValueString struct {
	Key string `json:"key,omitempty"`

	Val string `json:"val,omitempty"`
}

type KeyValueInt struct {
	Key string `json:"key,omitempty"`

	Val int `json:"val,omitempty"`
}
