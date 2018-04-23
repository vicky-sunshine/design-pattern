package duck

import (
	"fmt"
	"strategy/behavior"
)

type Duck interface {
	Swim()
	P()
	PerformFly()
}

type MallardDuck struct {
	quackBehavior behavior.QuackBehavior
	flyBehavior   behavior.FlyBehavior
}

func NewMallardDuck() *MallardDuck {
	duck := new(MallardDuck)
	duck.flyBehavior = behavior.NewFlyWithWings()
	duck.quackBehavior = behavior.NewQuack()

	return duck
}

func (d MallardDuck) swim() {
	fmt.Println("Swim!")
}

func (d MallardDuck) PerformQuack() {
	d.quackBehavior.Quack()
}

func (d MallardDuck) PerformFly() {
	d.flyBehavior.Fly()
}
