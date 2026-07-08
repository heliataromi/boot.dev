package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var marshalledData [][]byte

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		marshalledData = append(marshalledData, data)
	}

	return marshalledData, nil
}
