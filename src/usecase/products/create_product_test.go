package products

import (
	"context"
	"errors"
	"github.com/rasyidmm/superindo-test/src/adapter/repository/database/entity"
	"github.com/rasyidmm/superindo-test/src/shared/constanta"
	"github.com/rasyidmm/superindo-test/src/shared/mock"
	"github.com/rasyidmm/superindo-test/src/usecase/payload"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestProductsUsecase_CreateProduct(t *testing.T) {
	repo := &mock.ProductRepository{}
	redis := &mock.RedisRepository{}
	Convey("test create product", t, func() {
		Convey("Positive", func() {
			repo.On("CreateProduct", context.Background(), &entity.Product{Types: "BUAH", Price: 0, CreatedAt: time.Time{}}).Return(nil).Once()
			redis.On("DeleteCache", context.Background(), []string{constanta.ProductList}).Return(nil).Once()
			uc := NewProductUsecase(repo, redis)
			err := uc.CreateProduct(context.Background(), payload.CreateProductRequest{Types: "BUAH", Price: 0})
			So(err, ShouldBeNil)
		})
		Convey("Negative", func() {
			repo.On("CreateProduct", context.Background(), &entity.Product{Types: "BUAH", Price: 0, CreatedAt: time.Time{}}).Return(errors.New("error")).Once()
			redis.On("DeleteCache", context.Background(), []string{constanta.ProductList}).Return(nil).Once()
			uc := NewProductUsecase(repo, redis)
			err := uc.CreateProduct(context.Background(), payload.CreateProductRequest{Types: "BUAH", Price: 0})
			So(err, ShouldNotBeNil)
		})
	})
}
