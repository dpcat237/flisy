package plane

type Plane struct {
	MiddleSize int
	SideSize   int
	RowsNumber int
}

func (Plane) TableName() string {
	return "plane"
}
