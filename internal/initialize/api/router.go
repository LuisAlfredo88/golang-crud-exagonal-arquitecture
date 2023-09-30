package api

import (
	vehicle "golang-crud-exagonal-arquitecture/internal/modules/vehicle/infraestructure/api"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		vehicle.NewVehicleApi,
	),
)

func RegisterRoutes(
	s *API,
	e *echo.Echo,
	vehicleApi *vehicle.VehicleApi,
) {
	go s.Start(e)
	vehicleApi.Register()
}
