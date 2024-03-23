package products

import (
	"context"
	"github.com/rasyidmm/superindo-test/src/adapter/repository/database/entity"
	repoProductdb "github.com/rasyidmm/superindo-test/src/adapter/repository/database/products"
	repoRedis "github.com/rasyidmm/superindo-test/src/adapter/repository/redis"
	"github.com/rasyidmm/superindo-test/src/usecase/payload"
)

type ProductsUsecase struct {
	rpd repoProductdb.ProductRepository
	rpr repoRedis.RedisRepository
}

func NewProductUsecase(rpd repoProductdb.ProductRepository, rpr repoRedis.RedisRepository) *ProductsUsecase {
	return &ProductsUsecase{
		rpd: rpd,
		rpr: rpr,
	}
}

type ProductsPort interface {
	GetProducts(ctx context.Context, input payload.GetProductsRequest) (*[]entity.Product, error)
	CreateProduct(ctx context.Context, input payload.CreateProductRequest) error
}
