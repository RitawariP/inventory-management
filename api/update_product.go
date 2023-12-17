package api

import (
	"github.com/inventory-management/logic"
	"github.com/inventory-management/utils"
	"gofr.dev/pkg/gofr"
)

func UpdateProduct(ctx *gofr.Context) (interface{}, error) {
	requestBody, err := utils.ParseRequestBody(ctx.Request())
	if err != nil {
		return nil, err
	}

	product, err := utils.ValidateAndGetProduct(requestBody)
	if err != nil {
		return nil, err
	}

	err = logic.UpdateProduct(ctx, *product)
	if err != nil {
		return nil, err
	}
	return nil, err
}
