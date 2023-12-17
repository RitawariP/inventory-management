package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseRequestBody(r *http.Request) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading request body: %w", err)
	}
	defer r.Body.Close()

	// Convert JSON to map
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling request body: %w", err)
	}
	return data, nil
}
