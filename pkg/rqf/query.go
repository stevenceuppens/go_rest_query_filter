package rqf

import "encoding/json"

// Query ...
type Query struct {
	Fields map[string]bool        `json:"fields"`
	Where  map[string]interface{} `json:"where"`
	Order  []string               `json:"order"`
	Offset int                    `json:"offset"`
	Limit  int                    `json:"limit"`
}

// NewQuery ...
func NewQuery() *Query {
	query := &Query{
		Offset: 0,
		Limit:  -1,
	}
	return query
}

func (q *Query) String() string {
	json, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(json)
}
