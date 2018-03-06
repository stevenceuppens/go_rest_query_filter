package rqf

import (
	"encoding/json"
	"net/url"
	"strings"
)

// ParseFilter ...
func ParseFilter(filter string) (*Query, error) {

	normalizedFilter, err := normalizeFilter(filter)
	if err != nil {
		return nil, err
	}

	query, err := parseJSONFilter(normalizedFilter)
	if err != nil {
		return nil, err
	}

	return query, nil
}

// normalizeFilter will filter the 'filter' query parameter from the string
func normalizeFilter(filter string) (string, error) {
	stripped := filter

	index := strings.Index(stripped, "?filter=")
	if index != -1 {
		stripped = stripped[index+8:]
	}

	decoded, err := url.QueryUnescape(stripped)

	if err != nil {
		return "", err
	}

	return decoded, nil
}

// parseJSONFilter will parse the json into a Query object
func parseJSONFilter(filter string) (*Query, error) {
	query := NewQuery()

	err := json.Unmarshal([]byte(filter), query)
	if err != nil {
		return nil, err
	}

	return query, nil
}
