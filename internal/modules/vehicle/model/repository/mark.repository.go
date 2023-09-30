package repository

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
)

type MarkRepository interface {
	SaveMark(vehicle entity.Mark) error
	GetAllMarks() ([]entity.Mark, error)
	GetMark(markId int32) (entity.Mark, error)
	DeleteMark(mark int32) error

	SaveModel(vehicle entity.VehicleModel) error
	GetAllModels() ([]entity.VehicleModel, error)
	GetModel(markId int32) (entity.VehicleModel, error)
	DeleteModel(mark int32) error
}
