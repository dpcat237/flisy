package flight

import (
	"github.com/jinzhu/gorm"

	"gitlab.com/dpcat237/flisy/src/module/plane"
)

type Flight struct {
	gorm.Model

	PlaneID uint
	Plane   plane.Plane
}

func (Flight) TableName() string {
	return "flight"
}
