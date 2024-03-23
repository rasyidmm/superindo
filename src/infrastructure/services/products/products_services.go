package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasyidmm/superindo-test/src/usecase/payload"
	usecase "github.com/rasyidmm/superindo-test/src/usecase/products"
)

type ProductServices struct {
	uce usecase.ProductsPort
}

func NewProductServices(uc usecase.ProductsPort) *ProductServices {
	return &ProductServices{
		uce: uc,
	}
}

// GetProducts: Handler function for Get Products
// @Summary     GGet Products
// @Description Endpoint for  Get Products
// @Tags        Products
// @Products     json
// @Param       find    query     string  false  "find product" example(buah)
// @Param       sort  query     string  false  "sort" example(Buah)
// @Param       filter    query     string  false  "filter" example(created_at.desc)
// @Success     200     {object} []entity.Product
// @Router      /products [GET]
func (ps *ProductServices) GetProducts(ctx *fiber.Ctx) error {

	req := payload.GetProductsRequest{
		Find:   ctx.Query("find"),
		Sort:   ctx.Query("sort"),
		Filter: ctx.Query("filter"),
	}
	res, err := ps.uce.GetProducts(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

// CreateProduct: Handler function for Create Product
// @Summary    Create Product
// @Description Endpoint for Create Product
// @Tags        Products
// @Products     json
// @Param request body payload.CreateProductRequest true "request"
// @Success     200     {object} string
// @Router      /products [POST]
func (ps *ProductServices) CreateProduct(ctx *fiber.Ctx) error {
	var (
		req payload.CreateProductRequest
	)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := ps.uce.CreateProduct(ctx.Context(), req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}
