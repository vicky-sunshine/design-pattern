package behavior

import "fmt"

// FlyBehavior
type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct {
}

func NewFlyWithWings() *FlyWithWings {
	return new(FlyWithWings)
}

func (fww FlyWithWings) Fly() {
	fmt.Println("Fly with wings")
}

type FylNoWay struct {
}

func NewFylNoWay() *FylNoWay {
	return new(FylNoWay)
}

func (fnw FylNoWay) Fly() {
	fmt.Println("I can't fly")
}
