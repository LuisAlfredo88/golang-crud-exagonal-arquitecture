package api

import (
	"net/http"
	"strconv"

	shared "golang-crud-exagonal-arquitecture/internal/modules/shared/model"
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
	vehicle "golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/service"

	"github.com/labstack/echo/v4"
)

type VehicleApi struct {
	service     vehicle.VehicleService
	markService vehicle.MarkService
	api         *echo.Echo
}

func NewVehicleApi(
	service vehicle.VehicleService,
	markService vehicle.MarkService,
	api *echo.Echo,
) *VehicleApi {
	rest_api := &VehicleApi{
		service:     service,
		markService: markService,
		api:         api,
	}

	return rest_api
}

func (l *VehicleApi) Register() {
	vehicles := l.api.Group("/vehicles")
	vehicles.GET("", l.getAllVehicles)
	vehicles.POST("", l.saveVehicle)
	vehicles.PUT("", l.saveVehicle)
	vehicles.DELETE("/:id", l.deleteVehcile)

	marks := l.api.Group("/marks")
	marks.GET("", l.getAllMarks)
	marks.POST("", l.saveMark)
	marks.PUT("", l.saveMark)
	marks.DELETE("/:id", l.deleteMark)

	models := l.api.Group("/models")
	models.GET("", l.getAllModels)
	models.POST("", l.saveModel)
	models.PUT("", l.saveModel)
	models.DELETE("/:id", l.deleteModel)
}

func (p *VehicleApi) getAllVehicles(c echo.Context) error {
	vehicles, _ := p.service.GetAllVehicles()
	return c.JSON(http.StatusOK, vehicles)
}

func (p *VehicleApi) saveVehicle(c echo.Context) error {
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

func (p *VehicleApi) deleteVehcile(c echo.Context) error {
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
