package repository

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/infraestructure/repository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		repository.NewVehicleRepository,
		repository.NewMarkRepository,
	),
)
