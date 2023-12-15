package storage

import (
	"fmt"
	"github.com/inventory-management/model"
	"gofr.dev/pkg/gofr"
)

type ProductDao interface {
	CreateProduct(ctx *gofr.Context, product model.Product) error
	GetProductFromName(ctx *gofr.Context, name string) (*model.Product, error)
	GetAllProducts(ctx *gofr.Context) ([]model.Product, error)
	UpdateProduct(ctx *gofr.Context, product model.Product) error
}

type productDaoImpl struct{}

func NewProductDao() ProductDao {
	return &productDaoImpl{}
}

func (p *productDaoImpl) CreateProduct(ctx *gofr.Context, product model.Product) error {
	res, err := ctx.DB().ExecContext(ctx, insertQuery, product.Name, product.Description, product.Price, product.Created)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("insertion error: no rows affected")
	}
	return nil
}

func (p *productDaoImpl) GetProductFromName(ctx *gofr.Context, name string) (*model.Product, error) {
	var product model.Product
	err := ctx.DB().QueryRowContext(ctx, selectByNameQuery, name).
		Scan(&product.Name, &product.Description, &product.Price, &product.Created)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productDaoImpl) GetAllProducts(ctx *gofr.Context) ([]model.Product, error) {
	var products []model.Product
	rows, err := ctx.DB().QueryContext(ctx, selectAllQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.Name, &product.Description, &product.Price, &product.Created); err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func (p *productDaoImpl) UpdateProduct(ctx *gofr.Context, product model.Product) error {
	res, err := ctx.DB().ExecContext(ctx, updateQuery, product.Description, product.Price, product.Name)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("update error: no rows affected")
	}
	return nil
}
