package utils

import (
	"encoding/json"
	"fmt"
)

func ToSlice(obj interface{}) (map[string]interface{}, error) {
	// Marshal struct to JSON
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("marshal failed: %v", err)
	}

	// Unmarshal JSON to map
	var params map[string]interface{}
	if err := json.Unmarshal(data, &params); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %v", err)
	}

	return params, nil
}
