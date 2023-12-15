package server

import (
	"github.com/inventory-management/api"
	"gofr.dev/pkg/gofr"
)

func RegisterRoutes(app *gofr.Gofr) {
	app.GET("/product", api.GetProduct)
	app.POST("/product", api.CreateProduct)
	app.GET("/products", api.GetProducts)
}
