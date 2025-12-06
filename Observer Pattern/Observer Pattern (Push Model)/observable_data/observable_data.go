package observerable_data

type ObservableData struct {
	Count int
}

func NewObservableData(count int) *ObservableData {
	return &ObservableData{Count: count}
}

func (od *ObservableData) SetData(count int) {
	od.Count = count
}
