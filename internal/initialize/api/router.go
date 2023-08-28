package api

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(),
)

func RegisterRoutes(
	s *API,
	e *echo.Echo,
) {
	go s.Start(e)
}
