package products

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/rasyidmm/superindo-test/src/adapter/repository/database/entity"
	"github.com/rasyidmm/superindo-test/src/shared/constanta"
	"github.com/rasyidmm/superindo-test/src/usecase/payload"
	"time"
)

func (uc *ProductsUsecase) CreateProduct(ctx context.Context, input payload.CreateProductRequest) error {

	if !constanta.IsType[input.Types] {
		return fiber.NewError(fiber.StatusBadRequest, "type not found")
	}
	err := uc.rpd.CreateProduct(ctx, &entity.Product{
		SKU:       slug.Make(input.Name),
		Name:      input.Name,
		Types:     input.Types,
		Price:     input.Price,
		CreatedAt: time.Time{},
	})
	if err != nil {
		return err
	}

	uc.rpr.DeleteCache(ctx, constanta.ProductList)
	return nil
}
