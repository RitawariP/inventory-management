package api

import (
	"github.com/inventory-management/logic"
	"gofr.dev/pkg/gofr"
)

func GetProducts(ctx *gofr.Context) (interface{}, error) {
	products, err := logic.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
