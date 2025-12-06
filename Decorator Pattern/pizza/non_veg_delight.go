package pizza

type NonVegDelight struct {
}

func (nvd *NonVegDelight) GetPrice() float64 {
	return 7.0
}
