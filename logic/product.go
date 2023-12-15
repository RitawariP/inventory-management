package logic

import (
	"fmt"
	"github.com/inventory-management/model"
	"github.com/inventory-management/storage"
	"gofr.dev/pkg/gofr"
)

var productDao storage.ProductDao

func init() {
	productDao = storage.NewProductDao()
}

func CreateProduct(ctx *gofr.Context, product model.Product) error {
	saveErr := productDao.CreateProduct(ctx, product)
	if saveErr != nil {
		return fmt.Errorf("error saving product: %w", saveErr)
	}
	return nil
}
