package main

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/simuDuck/duckInterface"
	ducktypes "github.com/NandeeshG/designPatterns/simuDuck/duckTypes"
)

func main() {
	ducks := [...]duckInterface.DuckInterface{
		&ducktypes.MallardDuck{},
		&ducktypes.DecoyDuck{},
		&ducktypes.RubberDuck{},
	}
	for _, d := range ducks {
		fmt.Println("New duck upcoming!!")
		d.Init()
		d.PerformFly()
		d.PerformQuack()
		d.Display()
		d.Swim()
		fmt.Println()
	}
}
