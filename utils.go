package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func convertToReader(data interface{}) (io.Reader, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonData), nil
}
