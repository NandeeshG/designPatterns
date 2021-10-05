package ducktypes

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/fly"
	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/quack"
	"github.com/NandeeshG/designPatterns/simuDuck/duck"
)

type ModelDuck struct {
	duck.Duck
}

func (md *ModelDuck) Init() {
	md.SetFlyBehaviour(&fly.FlyNoWay{})
	md.SetQuackBehaviour(&quack.QuackMute{})
}

func (md *ModelDuck) Display() {
	fmt.Println("I am a model duck!")
}
