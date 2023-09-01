package repository

import (
	"golang-crud-exagonal-arquitecture/internal/modules/product/infraestructure/repository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		repository.NewLogRepository,
	),
)
