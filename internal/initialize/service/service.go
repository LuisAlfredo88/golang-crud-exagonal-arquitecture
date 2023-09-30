package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/application/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		service.NewVehicleService,
		service.NewMarkService,
	),
)
