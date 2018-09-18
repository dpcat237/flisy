package seat

import (
	"github.com/pkg/errors"

	"gitlab.com/dpcat237/flisy/src/module/flight"
	"gitlab.com/dpcat237/flisy/src/module/user"
)

type Handler interface {
	// Assign next available seat by type in order: aisle seats, window seats, middle seats.
	AssignSeat(u user.User, f flight.Flight) (SeatDTO, error)
	// GetSeat get specific seat information
	GetSeat(stIndex int) (SeatDTO, error)
	// GetSeats get ordered seats
	GetSeats(f flight.Flight) (SeatsDTO, error)
}

type handler struct {
	sRepo Repository
}

func NewHandler(sRepo Repository) *handler {
	return &handler{sRepo: sRepo}
}

func (stHnd *handler) AssignSeat(u user.User, f flight.Flight) (SeatDTO, error) {
	var sDTO SeatDTO

	nextSt, err := stHnd.getNextAvailable(f.ID)
	if err != nil {
		return sDTO, errors.New("All seats are assigned")
	}

	nextSt.UserID = u.ID
	if err := stHnd.sRepo.Save(&nextSt); err != nil {
		return sDTO, err
	}

	sDTO.LoadData(nextSt)
	return sDTO, nil
}

func (stHnd *handler) GetSeat(stIndex int) (SeatDTO, error) {
	var sDTO SeatDTO

	s, err := stHnd.sRepo.GetByIndex(stIndex)
	if err != nil {
		return sDTO, err
	}

	sDTO.LoadData(s)
	return sDTO, nil
}

func (stHnd *handler) GetSeats(f flight.Flight) (SeatsDTO, error) {
	var stsDTO SeatsDTO
	sts, err := stHnd.sRepo.GetByFlight(f.ID)
	if err != nil {
		return stsDTO, err
	}

	stsDTO.LoadData(sts, f.Plane)
	return stsDTO, nil
}

func (stHnd *handler) getNextAvailable(flightID uint) (Seat, error) {
	sa, _ := stHnd.sRepo.GetNextAvailable(flightID, typeAisle)
	if sa.ID > 0 {
		return sa, nil
	}

	sm, _ := stHnd.sRepo.GetNextAvailable(flightID, typeMiddle)
	if sm.ID > 0 {
		return sm, nil
	}

	return stHnd.sRepo.GetNextAvailable(flightID, typeWindow)
}
