package api

import (
	"fmt"
	"github.com/inventory-management/logic"
	"gofr.dev/pkg/gofr"
)

func DeleteProduct(ctx *gofr.Context) (interface{}, error) {
	productName := ctx.Params()["name"]
	if productName == "" {
		return nil, fmt.Errorf("product name is required")
	}

	err := logic.DeleteProduct(ctx, productName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
