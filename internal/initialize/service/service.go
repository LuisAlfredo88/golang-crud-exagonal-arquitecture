package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/application/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		service.NewProductService,
	),
)
