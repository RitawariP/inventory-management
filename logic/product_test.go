package logic

import (
	"fmt"
	"github.com/inventory-management/model"
	"github.com/inventory-management/storage"
	"github.com/inventory-management/storage/mocks"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	ctx := &gofr.Context{}
	type args struct {
		ctx     *gofr.Context
		product model.Product
	}
	tests := []struct {
		name    string
		args    args
		setup   func() storage.ProductDao
		wantErr error
	}{
		{
			name: "db error",
			args: args{
				ctx:     ctx,
				product: model.Product{},
			},
			setup: func() storage.ProductDao {
				mockProductDao := mocks.NewProductDao(t)
				mockProductDao.On("CreateProduct", ctx, model.Product{}).
					Return(fmt.Errorf("insertion error"))
				return mockProductDao
			},
			wantErr: fmt.Errorf("insertion error"),
		},
		{
			name: "success",
			args: args{
				ctx:     ctx,
				product: model.Product{Name: "dummy"},
			},
			setup: func() storage.ProductDao {
				mockProductDao := mocks.NewProductDao(t)
				mockProductDao.On("CreateProduct", ctx, model.Product{Name: "dummy"}).
					Return(nil)
				return mockProductDao
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productDao = tt.setup()
			err := CreateProduct(tt.args.ctx, tt.args.product)
			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
			}
		})
	}
}

func TestGetProductFromName(t *testing.T) {
	ctx := &gofr.Context{}
	product := model.Product{
		ID:          1,
		Name:        "iPhone",
		Description: "Apple iPhone",
		Price:       1000,
		Created:     "2021-01-01",
	}
	type args struct {
		ctx         *gofr.Context
		productName string
	}
	tests := []struct {
		name    string
		args    args
		setup   func() storage.ProductDao
		wantRes *model.Product
		wantErr error
	}{
		{
			name: "db error: product not found",
			args: args{
				ctx:         ctx,
				productName: "invalid",
			},
			setup: func() storage.ProductDao {
				mockProductDao := mocks.NewProductDao(t)
				mockProductDao.On("GetProductFromName", ctx, "invalid").
					Return(nil, fmt.Errorf("product not found"))
				return mockProductDao
			},
			wantRes: nil,
			wantErr: fmt.Errorf("product not found"),
		},
		{
			name: "success",
			args: args{
				ctx:         ctx,
				productName: "iPhone",
			},
			setup: func() storage.ProductDao {
				mockProductDao := mocks.NewProductDao(t)
				mockProductDao.On("GetProductFromName", ctx, "iPhone").
					Return(&product, nil)
				return mockProductDao
			},
			wantRes: &product,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productDao = tt.setup()
			res, err := GetProductFromName(tt.args.ctx, tt.args.productName)
			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
			}
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	ctx := &gofr.Context{}
	products := []model.Product{
		{
			ID:          1,
			Name:        "iPhone",
			Description: "Apple iPhone",
			Price:       1000,
			Created:     "2021-01-01",
		},
		{
			ID:          2,
			Name:        "iPad",
			Description: "Apple iPad",
			Price:       2000,
			Created:     "2021-01-01",
		},
	}
	type args struct {
		ctx *gofr.Context
	}
	tests := []struct {
		name    string
		args    args
		setup   func() storage.ProductDao
		wantRes []model.Product
		wantErr error
	}{
		{
			name: "db error",
			args: args{
				ctx: ctx,
			},
			setup: func() storage.ProductDao {
				mockProductDao := mocks.NewProductDao(t)
				mockProductDao.On("GetAllProducts", ctx).
					Return(nil, fmt.Errorf("internal error"))
				return mockProductDao
			},
			wantRes: nil,
			wantErr: fmt.Errorf("internal error"),
		},
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			setup: func() storage.ProductDao {
				mockProductDao := mocks.NewProductDao(t)
				mockProductDao.On("GetAllProducts", ctx).
					Return(products, nil)
				return mockProductDao
			},
			wantRes: products,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productDao = tt.setup()
			res, err := GetAllProducts(tt.args.ctx)
			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
			}
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
