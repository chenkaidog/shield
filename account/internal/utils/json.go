package utils

import (
	"encoding/json"
)

func SafeJson(e interface{}) string {
	data, err := json.Marshal(e)
	if err != nil {
		return "{}"
	}

	return string(data)
}
