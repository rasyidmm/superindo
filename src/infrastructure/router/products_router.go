package router

import (
	"github.com/gofiber/fiber/v2"
	service "github.com/rasyidmm/superindo-test/src/infrastructure/services/products"
)

func NewProducts(f fiber.Router, service *service.ProductServices) *fiber.Router {
	r := f.Group("/products")
	r.Get("", service.GetProducts)
	r.Post("", service.CreateProduct)
	return &f
}
