package entity

type Mark struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
}
