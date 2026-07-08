package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &resources); err != nil {
	    return nil, err
	}
	
	return resources, nil
}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for _, resource := range resources {
		for key, value := range resource {
			formattedString := fmt.Sprintf("Key: %s - Value: %v", key, value)
			formattedStrings = append(formattedStrings, formattedString)
		}
	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}
