package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
)

type MarkService interface {
	SaveMark(mark entity.Mark) error
	GetAllMarks() ([]entity.Mark, error)
	GetMark(markId int32) (entity.Mark, error)
	DeleteMark(mark int32) error

	SaveModel(mark entity.VehicleModel) error
	GetAllModels() ([]entity.VehicleModel, error)
	GetModel(markId int32) (entity.VehicleModel, error)
	DeleteModel(mark int32) error
}
