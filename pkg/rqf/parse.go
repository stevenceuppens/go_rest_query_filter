package rqf

import (
	"encoding/json"
	"net/url"
	"strings"
)

// ParseFilter ...
func ParseFilter(data string) (*Filter, error) {

	normalizedFilter, err := normalizeFilter(data)
	if err != nil {
		return nil, err
	}

	filter, err := parseJSONFilter(normalizedFilter)
	if err != nil {
		return nil, err
	}

	return filter, nil
}

// normalizeFilter will filter the 'filter' query parameter from the string
func normalizeFilter(data string) (string, error) {
	stripped := data

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
func parseJSONFilter(data string) (*Filter, error) {
	filter := NewFilter()

	err := json.Unmarshal([]byte(data), filter)
	if err != nil {
		return nil, err
	}

	return filter, nil
}
