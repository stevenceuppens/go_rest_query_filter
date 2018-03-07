package rqf

import "encoding/json"

// Filter ...
type Filter struct {
	Fields map[string]bool        `json:"fields"`
	Where  map[string]interface{} `json:"where"`
	Order  []string               `json:"order"`
	Offset int                    `json:"offset"`
	Limit  int                    `json:"limit"`
}

// NewFilter ...
func NewFilter() *Filter {
	filter := &Filter{
		Offset: 0,
		Limit:  0,
	}
	return filter
}

// String ...
func (q *Filter) String() string {
	json, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(json)
}
