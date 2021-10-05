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
	d.FlyBehaviour = &fly.FlyNoWay{}
	d.QuackBehaviour = &quack.QuackSqueak{}
}

func (d *RubberDuck) Display() {
	fmt.Println("I am a rubber ducky!")
}
