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

func TestProductsUsecase_GetProducts(t *testing.T) {
	repo := &mock.ProductRepository{}
	redis := &mock.RedisRepository{}

	Convey("test get products", t, func() {
		Convey("Positive", func() {
			repo.On("GetProducts", context.Background()).Return(&[]entity.Product{}, nil).Once()
			redis.On("GetCache", context.Background(), constanta.ProductList).Return("", nil).Once()
			redis.On("SetCache", context.Background(), constanta.ProductList, "[]", 12*time.Hour).Return(nil).Once()
			uc := NewProductUsecase(repo, redis)
			data, err := uc.GetProducts(context.Background(), payload.GetProductsRequest{})
			So(err, ShouldBeNil)
			So(data, ShouldNotBeNil)
		})
		Convey("Negative", func() {
			repo.On("GetProducts", context.Background()).Return(nil, errors.New("error")).Once()
			redis.On("GetCache", context.Background(), constanta.ProductList).Return("", nil).Once()
			redis.On("SetCache", context.Background(), constanta.ProductList, "null", 12*time.Hour).Return(nil).Once()
			uc := NewProductUsecase(repo, redis)
			data, err := uc.GetProducts(context.Background(), payload.GetProductsRequest{})
			So(err, ShouldNotBeNil)
			So(data, ShouldBeNil)
		})
	})
}
