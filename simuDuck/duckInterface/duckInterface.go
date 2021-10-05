package duckInterface

import (
	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/fly"
	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/quack"
)

type DuckInterface interface {
	Init()
	SetFlyBehaviour(fly.FlyBehaviour)
	SetQuackBehaviour(quack.QuackBehaviour)
	PerformFly()
	PerformQuack()
	Display()
	Swim()
}
