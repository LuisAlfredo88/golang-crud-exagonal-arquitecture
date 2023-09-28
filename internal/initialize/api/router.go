package api

import (
	vehcile "golang-crud-exagonal-arquitecture/internal/modules/vehicle/infraestructure/api"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		vehcile.NewVehcileApi,
	),
)

func RegisterRoutes(
	s *API,
	e *echo.Echo,
	vehicleApi *vehcile.VehcileApi,
) {
	go s.Start(e)
	vehicleApi.Register()
}
