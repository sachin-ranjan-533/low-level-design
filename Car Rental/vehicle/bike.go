package vehicle

type Bike struct {
	Id   int
	Name string
}

func NewBike(id int, name string) *Bike {
	return &Bike{
		Id:   id,
		Name: name,
	}
}
