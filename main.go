package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/rasyidmm/superindo-test/docs"
	"github.com/rasyidmm/superindo-test/internal/config"
	"github.com/rasyidmm/superindo-test/src/infrastructure/router"
	serviceProduct "github.com/rasyidmm/superindo-test/src/infrastructure/services/products"
	container "github.com/rasyidmm/superindo-test/src/shared/di"
	"github.com/rasyidmm/superindo-test/src/usecase/products"
)

// @title           superindo
// @version         1.0
// @description     This is a backend service for superindo
// @termsOfService  http://swagger.io/terms/

// @contact.name   rasyid
// @contact.email  rasyidmaulidmajid@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath
func main() {
	app := fiber.New(
		fiber.Config{
			AppName: config.GetConfig().Server.Rest.Host,
		})

	app.Get("/swagger/*", swagger.HandlerDefault)
	ctn := container.NewContainer()
	apply(app, ctn)
	err := app.Listen(":" + string(config.GetConfig().Server.Rest.Port))
	if err != nil {
		panic(err)
	}
}

func apply(f fiber.Router, ctn *container.Container) {
	router.NewProducts(f, serviceProduct.NewProductServices(ctn.Get("product").(*products.ProductsUsecase)))
}
