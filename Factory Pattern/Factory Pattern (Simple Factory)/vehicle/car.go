package vehicle

import "fmt"

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Driving a car")
}
