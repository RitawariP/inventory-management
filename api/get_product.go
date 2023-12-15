package api

import (
	"fmt"
	"github.com/inventory-management/logic"
	"gofr.dev/pkg/gofr"
)

func GetProduct(ctx *gofr.Context) (interface{}, error) {
	productName := ctx.Params()["name"]
	if productName == "" {
		return nil, fmt.Errorf("product name is required")
	}

	product, err := logic.GetProductFromName(ctx, productName)
	if err != nil {
		return nil, err
	}
	return product, nil
}
