package entity

import (
	"errors"
	"strings"
)

func ValidateVehicle(vehicle Vehicle) error {
	if strings.TrimSpace(vehicle.Name) == "" {
		return errors.New("must specify vehicle name")
	}

	return nil
}
