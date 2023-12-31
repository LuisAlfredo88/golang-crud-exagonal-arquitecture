package entity

import "gorm.io/gorm"

type VehicleModel struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name"`
	MarkID uint   `json:"mark_id"`
	Mark   Mark
}
