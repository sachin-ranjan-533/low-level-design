package adaptee

type WeightMachineBaby struct {
	weightInPounds float64
}

func NewWeightMachineBaby(weight float64) *WeightMachineBaby {
	return &WeightMachineBaby{
		weightInPounds: weight,
	}
}

func (wmb *WeightMachineBaby) GetWeightInPounds() float64 {
	return wmb.weightInPounds
}
