package di

import (
	"github.com/rasyidmm/superindo-test/src/adapter/connection"
	repoProductdb "github.com/rasyidmm/superindo-test/src/adapter/repository/database/products"
	repoRedis "github.com/rasyidmm/superindo-test/src/adapter/repository/redis"
	"github.com/rasyidmm/superindo-test/src/usecase/products"
	"github.com/sarulabs/di/v2"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "product", Build: productsUsecase},
	}...)
	return &Container{ctn: builder.Build()}
}

func (c *Container) Get(name string) interface{} {
	return c.ctn.Get(name)
}

func productsUsecase(_ di.Container) (interface{}, error) {
	repodb := repoProductdb.NewProductsDataHandler(connection.DB)
	repoRedis := repoRedis.NewRedisDataHandler(connection.ClientRedis)
	return products.NewProductUsecase(repodb, repoRedis), nil

}
