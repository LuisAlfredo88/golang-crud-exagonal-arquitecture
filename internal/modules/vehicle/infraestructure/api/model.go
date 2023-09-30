package api

import (
	"net/http"
	"strconv"

	shared "golang-crud-exagonal-arquitecture/internal/modules/shared/model"
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"

	"github.com/labstack/echo/v4"
)

func (p *VehicleApi) getAllModels(c echo.Context) error {
	models, _ := p.markService.GetAllModels()
	return c.JSON(http.StatusOK, models)
}

func (p *VehicleApi) saveModel(c echo.Context) error {
	params := entity.VehicleModel{}
	c.Bind(&params)

	err := p.markService.SaveModel(params)

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, nil)
}

func (p *VehicleApi) deleteModel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	p.markService.DeleteModel(int32(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, nil)
}
