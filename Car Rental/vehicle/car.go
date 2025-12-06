package vehicle

type Car struct {
	Id   int
	Name string
}

func NewCar(id int, name string) *Car {
	return &Car{
		Id:   id,
		Name: name,
	}
}
