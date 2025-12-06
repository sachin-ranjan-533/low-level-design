package adapter

import "adapter-pattern/adaptee"

type Adapter struct {
	weightMachine adaptee.WeightMachine
}

func NewAdapter(wmb adaptee.WeightMachine) *Adapter {
	return &Adapter{
		weightMachine: wmb,
	}
}

func (a *Adapter) GetWeightInKg() float64 {
	const poundsToKgFactor = 0.453592
	weightInPounds := a.weightMachine.GetWeightInPounds()
	return weightInPounds * poundsToKgFactor
}
