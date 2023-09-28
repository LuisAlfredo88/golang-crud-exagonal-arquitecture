package api

import (
	"net/http"
	"strconv"

	shared "golang-crud-exagonal-arquitecture/internal/modules/shared/model"
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
	vehicle "golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/service"

	"github.com/labstack/echo/v4"
)

type VehcileApi struct {
	service vehicle.VehicleService
	api     *echo.Echo
}

func NewVehcileApi(
	service vehicle.VehicleService,
	api *echo.Echo,
) *VehcileApi {
	rest_api := &VehcileApi{
		service: service,
		api:     api,
	}

	return rest_api
}

func (l *VehcileApi) Register() {
	logs := l.api.Group("/vehicles")
	logs.GET("", l.getAllVehicles)
	logs.POST("", l.saveVehicle)
	logs.PUT("", l.saveVehicle)
	logs.DELETE("/:id", l.deleteVehcile)
}

func (p *VehcileApi) getAllVehicles(c echo.Context) error {
	vehicles, _ := p.service.GetAllVehicles()
	return c.JSON(http.StatusOK, vehicles)
}

func (p *VehcileApi) saveVehicle(c echo.Context) error {
	params := entity.Vehicle{}
	c.Bind(&params)

	err := p.service.SaveVehicle(params)

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, nil)
}

func (p *VehcileApi) deleteVehcile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	p.service.DeleteVehicle(int32(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, nil)
}
