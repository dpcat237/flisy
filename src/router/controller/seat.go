package controller

import (
	"net/http"
	"strconv"

	"gitlab.com/dpcat237/flisy/src/module/flight"
	"gitlab.com/dpcat237/flisy/src/module/seat"
)

type seatController struct {
	fHnd  flight.Handler
	stHnd seat.Handler
}

func newSeatController(fHnd flight.Handler, stHnd seat.Handler) *seatController {
	return &seatController{fHnd: fHnd, stHnd: stHnd}
}

func (stC *seatController) AssignSeat(w http.ResponseWriter, r *http.Request) {
	u, err := getUser(r) // For demo propose this user was manually set
	if err != nil {
		returnPreconditionFailed(w, "")
	}

	flightID := uint(1) // For demo propose manually set
	f, err := stC.fHnd.GetByID(flightID)
	if err != nil {
		returnPreconditionFailed(w, err.Error())
		return
	}

	st, err := stC.stHnd.AssignSeat(u, f)
	if err != nil {
		returnErrorWithStatus(w, err.Error(), http.StatusNotFound)
	}
	returnJson(w, st)
}

func (stC *seatController) GetSeat(w http.ResponseWriter, r *http.Request) {
	stIndex, err := strconv.Atoi(getVariable(r, "index"))
	if err != nil || stIndex < 1 {
		returnPreconditionFailed(w, "")
	}

	st, err := stC.stHnd.GetSeat(stIndex)
	if err != nil {
		returnServerFailed(w, err.Error())
		return
	}
	returnJson(w, st)
}

func (stC *seatController) GetSeats(w http.ResponseWriter, r *http.Request) {
	flightID := uint(1) // For demo propose manually set
	f, err := stC.fHnd.GetByID(flightID)
	if err != nil {
		returnPreconditionFailed(w, err.Error())
		return
	}

	seats, err := stC.stHnd.GetSeats(f)
	if err != nil {
		returnServerFailed(w, err.Error())
		return
	}
	returnJson(w, seats)
}
