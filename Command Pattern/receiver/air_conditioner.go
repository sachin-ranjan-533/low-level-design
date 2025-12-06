package receiver

import "fmt"

type AirConditioner struct {
	brand       string
	temparature int
	isOn        bool
}

func NewAirConditioner(brand string) *AirConditioner {
	return &AirConditioner{
		brand:       brand,
		temparature: 24,
		isOn:        false,
	}
}

func (ac *AirConditioner) On() {
	fmt.Println("Air Conditioner is ON")
	ac.isOn = true
}

func (ac *AirConditioner) Off() {
	fmt.Println("Air Conditioner is OFF")
	ac.isOn = false
}

func (ac *AirConditioner) SetTemperature(temp int) {
	ac.temparature = temp
}

func (ac *AirConditioner) GetStatus() (bool, int) {
	return ac.isOn, ac.temparature
}
