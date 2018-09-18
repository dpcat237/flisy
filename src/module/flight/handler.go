package flight

type Handler interface {
	GetByID(flightID uint) (Flight, error)
}

type handler struct {
	fRepo Repository
}

func NewHandler(fRepo Repository) *handler {
	return &handler{fRepo: fRepo}
}

func (fHnd *handler) GetByID(flightID uint) (Flight, error) {
	return fHnd.fRepo.GetByID(flightID)
}
