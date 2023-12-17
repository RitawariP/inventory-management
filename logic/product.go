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

func GetProductFromName(ctx *gofr.Context, name string) (*model.Product, error) {
	product, err := productDao.GetProductFromName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("error getting product: %w", err)
	}
	return product, nil
}

func GetAllProducts(ctx *gofr.Context) ([]model.Product, error) {
	products, err := productDao.GetAllProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}
	return products, nil
}

func UpdateProduct(ctx *gofr.Context, product model.Product) error {
	existingProduct, err := productDao.GetProductFromName(ctx, product.Name)
	if err != nil {
		return fmt.Errorf("error getting existing product: %w", err)
	}

	product.ID = existingProduct.ID
	saveErr := productDao.UpdateProduct(ctx, product)
	if saveErr != nil {
		return fmt.Errorf("error updating product: %w", saveErr)
	}
	return nil
}

func DeleteProduct(ctx *gofr.Context, name string) error {
	existingProduct, err := productDao.GetProductFromName(ctx, name)
	if err != nil {
		return fmt.Errorf("error getting existing product: %w", err)
	}

	deleteErr := productDao.DeleteProduct(ctx, existingProduct.ID)
	if deleteErr != nil {
		return fmt.Errorf("error deleting product: %w", deleteErr)
	}
	return nil
}
