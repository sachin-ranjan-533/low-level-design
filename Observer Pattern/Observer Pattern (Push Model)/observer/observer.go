package observer

import observerable_data "observable-pattern/observable_data"

type ObservableData struct {
	count int
}

type Observer interface {
	Update(observableData *observerable_data.ObservableData)
}
