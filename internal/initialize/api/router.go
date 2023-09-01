package api

import (
	product "golang-crud-exagonal-arquitecture/internal/modules/product/infraestructure/api"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		product.NewProductApi,
	),
)

func RegisterRoutes(
	s *API,
	e *echo.Echo,
	productApy *product.ProductApi,
) {
	go s.Start(e)
	productApy.Register()
}
