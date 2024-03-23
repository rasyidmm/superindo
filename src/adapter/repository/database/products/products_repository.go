package products

import (
	"context"
	"github.com/rasyidmm/superindo-test/src/adapter/repository/database/entity"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, input *entity.Product) error
	GetProducts(ctx context.Context) (*[]entity.Product, error)
}
