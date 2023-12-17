package utils

import (
	"fmt"
	"github.com/inventory-management/model"
	"time"
)

func ValidateAndGetProduct(body map[string]interface{}) (*model.Product, error) {
	name, ok := body["name"].(string)
	if !ok {
		return nil, fmt.Errorf("name is required")
	}
	description, ok := body["description"].(string)
	if !ok {
		return nil, fmt.Errorf("description is required")
	}
	price, ok := body["price"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid price")
	}

	return &model.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Created:     time.Now().Format("2006-01-02"),
	}, nil
}
