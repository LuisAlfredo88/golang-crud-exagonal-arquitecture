package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type API struct {
}

func NewHTTPServer(lc fx.Lifecycle, srv *http.ServeMux) *API {
	return &API{}
}

func (s *API) Start(e *echo.Echo) {
	e.Start(":3000")
}
