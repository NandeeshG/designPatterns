package ducktypes

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/fly"
	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/quack"
	"github.com/NandeeshG/designPatterns/simuDuck/duck"
)

type RubberDuck struct {
	duck.Duck
}

func (d *RubberDuck) Init() {
	d.SetFlyBehaviour(&fly.FlyNoWay{})
	d.SetQuackBehaviour(&quack.QuackSqueak{})
}

func (d *RubberDuck) Display() {
	fmt.Println("I am a rubber ducky!")
}
