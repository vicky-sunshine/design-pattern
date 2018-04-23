package main

import (
	"strategy/duck"
)

func main() {
	d := duck.NewMallardDuck()
	d.PerformFly()
	d.PerformQuack()
}
