package behavior

import (
	"fmt"
)

// QuackBehavior
type QuackBehavior interface {
	Quack()
}

type Quack struct {
}

func NewQuack() *Quack {
	return new(Quack)
}

func (q Quack) Quack() {
	fmt.Println("Quack!")
}

type Queak struct {
}

func NewQeack() *Queak {
	return new(Queak)
}

func (q Queak) Quack() {
	fmt.Println("queak!")
}

type MuteQuack struct {
}

func NewMuteQuack() *MuteQuack {
	return new(MuteQuack)
}

func (q MuteQuack) Quack() {
	fmt.Println("<silence>")
}
