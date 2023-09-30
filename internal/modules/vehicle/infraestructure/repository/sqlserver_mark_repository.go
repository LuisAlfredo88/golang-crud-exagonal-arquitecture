package repository

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
	markModel "golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/repository"

	"gorm.io/gorm"
)

type markRepo struct {
	db *gorm.DB
}

func NewMarkRepository(db *gorm.DB) markModel.MarkRepository {
	return &markRepo{
		db: db,
	}
}

func (r *markRepo) SaveMark(markEntity entity.Mark) error {
	result := r.db.Where("id = ?", markEntity.ID).Save(&markEntity)
	return result.Error
}

func (r *markRepo) GetAllMarks() ([]entity.Mark, error) {
	var marks []entity.Mark
	r.db.Find(&marks)
	return marks, nil
}

func (r *markRepo) GetMark(markId int32) (entity.Mark, error) {
	var mark entity.Mark
	r.db.Where("id = ?", markId).Find(&mark)
	return mark, nil
}

func (r *markRepo) DeleteMark(markId int32) error {
	result := r.db.Delete(&entity.Mark{}, markId)
	return result.Error
}

func (r *markRepo) SaveModel(markEntity entity.VehicleModel) error {
	result := r.db.Where("id = ?", markEntity.ID).Save(&markEntity)
	return result.Error
}

func (r *markRepo) GetAllModels() ([]entity.VehicleModel, error) {
	var models []entity.VehicleModel
	r.db.Find(&models)
	return models, nil
}

func (r *markRepo) GetModel(modelId int32) (entity.VehicleModel, error) {
	var model entity.VehicleModel
	r.db.Where("id = ?", modelId).Find(&model)
	return model, nil
}

func (r *markRepo) DeleteModel(modelId int32) error {
	result := r.db.Delete(&entity.VehicleModel{}, modelId)
	return result.Error
}
