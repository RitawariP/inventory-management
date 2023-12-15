package storage

import (
	"fmt"
	"github.com/inventory-management/model"
	"gofr.dev/pkg/gofr"
)

type ProductDao interface {
	CreateProduct(ctx *gofr.Context, product model.Product) error
}

type productDaoImpl struct{}

func NewProductDao() ProductDao {
	return &productDaoImpl{}
}

func (p *productDaoImpl) CreateProduct(ctx *gofr.Context, product model.Product) error {
	res, err := ctx.DB().ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Created)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("insertion error: no rows affected")
	}
	return nil
}
