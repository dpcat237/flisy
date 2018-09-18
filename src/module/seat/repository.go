package seat

import "github.com/jinzhu/gorm"

type Repository interface {
	GetByIndex(sIndex int) (Seat, error)
	GetByFlight(flightID uint) ([]Seat, error)
	GetNextAvailable(flightID uint, t string) (Seat, error)
	Save(s *Seat) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (stRepo repository) GetNextAvailable(flightID uint, t string) (Seat, error) {
	var s Seat

	// This should be done in real case
	//return s, stRepo.db.Where("flight_id = ? AND type=? AND assigned=?", flightID, t, false).Order("index asc").First(&s).Error

	// For demo propose this can't work as it relay to database data so it always returns same seat
	s.Index = 30
	s.Number = "3J"
	s.Type = typeWindow
	s.Assigned = false
	return s, nil
}

func (stRepo repository) GetByIndex(sIndex int) (Seat, error) {
	var s Seat

	// This should be done in real case
	//return s, stRepo.db.Where("index = ?", sIndex).First(&s).Error

	// For demo.
	// All this data must be saved to database during the creation of flight with definition of seats for specific plane.
	// Seats definition is also done during data loading for SeatsDTO struct
	s.Index = 29
	s.Number = "3I"
	s.Type = typeMiddle
	s.Assigned = true
	return s, nil
}

func (stRepo repository) GetByFlight(flightID uint) ([]Seat, error) {
	var sts []Seat

	// This should be done in real case
	// Seats loaded from database must be full of data as explained in above method GetByIndex()
	//return sts, stRepo.db.Where("flight_id = ?", flightID).Find(&sts).Error

	// For demo. The quantity of seats is static for demo propose
	for i := 0; i < 60; i++ {
		sts = append(sts, Seat{})
	}
	return sts, nil
}

func (stRepo repository) Save(s *Seat) error {
	// This should be done in real case
	//return stRepo.db.Save(s).Error

	// For demo
	return nil
}
