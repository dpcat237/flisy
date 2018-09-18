package flight

import (
	"github.com/jinzhu/gorm"

	"gitlab.com/dpcat237/flisy/src/module/plane"
)

type Repository interface {
	GetByID(flightID uint) (Flight, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (fRepo repository) GetByID(flightID uint) (Flight, error) {
	var f Flight
	// This should be done in real case
	//return f, fRepo.db.Where("id = ?", flightID).Preload("Plane").First(&f).Error

	// For demo propose
	f.ID = flightID
	f.Plane = plane.Plane{
		MiddleSize: 4,
		SideSize:   3,
		RowsNumber: 60,
	}
	return f, nil
}
