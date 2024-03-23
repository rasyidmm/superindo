package mock

import (
	"context"
	"github.com/rasyidmm/superindo-test/src/adapter/repository/database/entity"
	"github.com/stretchr/testify/mock"
)

type ProductRepository struct {
	mock.Mock
}

func (m *ProductRepository) CreateProduct(ctx context.Context, input *entity.Product) error {
	args := m.Called(ctx, input)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (m *ProductRepository) GetProducts(ctx context.Context) (*[]entity.Product, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entity.Product), nil
}
