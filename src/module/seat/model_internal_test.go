package seat

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/dpcat237/flisy/src/module/plane"
)

func TestHandler_SeatsDTO_LoadData(t *testing.T) {
	var stsDTO SeatsDTO
	p := plane.Plane{
		MiddleSize: 4,
		SideSize:   3,
		RowsNumber: 60,
	}
	var sts []Seat
	for i := 0; i < 60; i++ {
		sts = append(sts, Seat{})
	}
	stsDTO.LoadData(sts, p)

	assert.Equal(t, 1, stsDTO.Rows[0].Seats[0].Index)
	assert.Equal(t, "1A", stsDTO.Rows[0].Seats[0].Number)
	assert.Equal(t, typeWindow, stsDTO.Rows[0].Seats[0].Type)
	assert.Equal(t, 29, stsDTO.Rows[2].Seats[8].Index)
	assert.Equal(t, "3I", stsDTO.Rows[2].Seats[8].Number)
	assert.Equal(t, typeMiddle, stsDTO.Rows[2].Seats[8].Type)
	assert.Equal(t, 582, stsDTO.Rows[58].Seats[2].Index)
	assert.Equal(t, "59C", stsDTO.Rows[58].Seats[2].Number)
	assert.Equal(t, typeAisle, stsDTO.Rows[58].Seats[2].Type)
}
