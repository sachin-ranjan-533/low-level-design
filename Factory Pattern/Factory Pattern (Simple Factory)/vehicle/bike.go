package vehicle

import "fmt"

type Bike struct{}

func (b *Bike) Drive() {
	fmt.Println("Riding a bike")
}
