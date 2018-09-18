package seat

import (
	"strconv"

	"github.com/jinzhu/gorm"

	"gitlab.com/dpcat237/flisy/src/module/plane"
)

const (
	typeAisle  = "aisle"
	typeMiddle = "middle"
	typeWindow = "window"
)

type Seat struct {
	gorm.Model

	FlightID uint
	Index    int
	Number   string
	Type     string
	Assigned bool
	UserID   uint
}

type SeatDTO struct {
	Index    int    `json:"index"`
	Number   string `json:"number"`
	Type     string `json:"type"`
	Assigned bool   `json:"assigned"`
}

type SeatsDTO struct {
	Rows []SeatsRowDTO `json:"rows"`
}

type SeatsRowDTO struct {
	Seats []SeatDTO `json:"seats"`
}

func (Seat) TableName() string {
	return "seat"
}

func (st *SeatDTO) LoadData(seat Seat) {
	st.Index = seat.Index
	st.Number = seat.Number
	st.Type = seat.Type
	st.Assigned = seat.Assigned
}

func (sr *SeatsRowDTO) CreateSeatDTO(i int, n, t string, a bool) SeatDTO {
	return SeatDTO{
		Index:    i,
		Number:   n,
		Type:     t,
		Assigned: a,
	}
}

func (sr *SeatsRowDTO) GetSeatNumber(ir, is int) string {
	return strconv.Itoa(ir) + string(rune(is+65))
}

func (sr *SeatsRowDTO) GetSeatType(c, max int, lTxt, mTxt, rTxt string) string {
	if c == 0 {
		return lTxt
	}
	if c == max-1 {
		return rTxt
	}
	return mTxt
}

// LoadData order seats by rows. For demo propose it also assigns to seat their Index, Number and Type
func (sts *SeatsDTO) LoadData(seats []Seat, p plane.Plane) {
	i := 1
	ir := 1
	for row := 1; row <= p.RowsNumber; row++ {
		var r SeatsRowDTO
		is := 0
		for s := 0; s < p.SideSize; s++ {
			r.Seats = append(r.Seats, r.CreateSeatDTO(i, r.GetSeatNumber(ir, is), r.GetSeatType(s, p.SideSize, typeWindow, typeMiddle, typeAisle), false))
			i++
			is++
		}
		for m := 0; m < p.MiddleSize; m++ {
			r.Seats = append(r.Seats, r.CreateSeatDTO(i, r.GetSeatNumber(ir, is), r.GetSeatType(m, p.MiddleSize, typeAisle, typeMiddle, typeAisle), false))
			i++
			is++
		}
		for s := 0; s < p.SideSize; s++ {
			r.Seats = append(r.Seats, r.CreateSeatDTO(i, r.GetSeatNumber(ir, is), r.GetSeatType(s, p.SideSize, typeAisle, typeMiddle, typeWindow), false))
			i++
			is++
		}
		sts.Rows = append(sts.Rows, r)
		ir++
	}
}
