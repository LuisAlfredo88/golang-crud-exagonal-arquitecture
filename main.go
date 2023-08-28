package main

import (
	"golang-crud-exagonal-arquitecture/internal/initialize/api"
	"golang-crud-exagonal-arquitecture/internal/initialize/db/sqlite"
	"golang-crud-exagonal-arquitecture/internal/initialize/repository"
	"golang-crud-exagonal-arquitecture/internal/initialize/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			echo.New,
			http.NewServeMux,
			sqlite.NewSQLite,
			api.NewHTTPServer,
		),
		repository.Module,
		service.Module,
		api.Module,
		fx.Invoke(api.RegisterRoutes),
	)

	app.Run()
}
