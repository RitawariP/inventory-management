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
