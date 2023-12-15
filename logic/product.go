package logic

import (
	"fmt"
	"github.com/inventory-management/model"
	"gofr.dev/pkg/gofr"
)

const query = "INSERT INTO products (name, description, price, created) VALUES (?, ?, ?, ?)"

func CreateProduct(ctx *gofr.Context, product model.Product) error {
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
