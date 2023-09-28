package entity

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	ID             uint    `gorm:"primaryKey;autoIncrement"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	Color          string  `json:"color"`
	Year           int32   `json:"year"`
	Chassis        string  `json:"chassis"`
	Plate          string  `json:"plate"`
	PassengersQty  string  `json:"passengers_qty"`
	DoorsQty       string  `json:"doors_qty"`
	Cylinders      string  `json:"cylinders"`
	Status         int8    `json:"status"`
	VehicleModelID uint
	VehicleModel   VehicleModel
}
