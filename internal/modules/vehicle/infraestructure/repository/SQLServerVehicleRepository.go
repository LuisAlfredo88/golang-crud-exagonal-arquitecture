package repository

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
	vehicleModel "golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/repository"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) vehicleModel.VehicleRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) SaveVehicle(vehicleEntity entity.Vehicle) error {
	result := r.db.Where("id = ?", vehicleEntity.ID).Save(&vehicleEntity)
	return result.Error
}

func (r *repo) GetAllVehicles() ([]entity.Vehicle, error) {
	var vehciles []entity.Vehicle
	r.db.Find(&vehciles)
	return vehciles, nil
}

func (r *repo) GetVehicle(vehicleId int32) (entity.Vehicle, error) {
	var vehicle entity.Vehicle
	r.db.Where("id = ?", vehicleId).Find(&vehicle)
	return vehicle, nil
}

func (r *repo) DeleteVehicle(vehcileId int32) error {
	result := r.db.Delete(&entity.Vehicle{}, vehcileId)
	return result.Error
}
