package api

import (
	"net/http"
	"strconv"

	shared "golang-crud-exagonal-arquitecture/internal/modules/shared/model"
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"

	"github.com/labstack/echo/v4"
)

func (p *VehicleApi) getAllMarks(c echo.Context) error {
	marks, _ := p.markService.GetAllMarks()
	return c.JSON(http.StatusOK, marks)
}

func (p *VehicleApi) saveMark(c echo.Context) error {
	params := entity.Mark{}
	c.Bind(&params)

	err := p.markService.SaveMark(params)

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, nil)
}

func (p *VehicleApi) deleteMark(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	p.markService.DeleteMark(int32(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, nil)
}
