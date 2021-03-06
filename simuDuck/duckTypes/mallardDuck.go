package ducktypes

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/fly"
	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/quack"
	"github.com/NandeeshG/designPatterns/simuDuck/duck"
)

type MallardDuck struct {
	duck.Duck
}

func (md *MallardDuck) Init() {
	md.SetFlyBehaviour(&fly.FlyWithWings{})
	md.SetQuackBehaviour(&quack.QuackNormal{})
}

func (md *MallardDuck) Display() {
	fmt.Println("I am a beautiful mallard duck!")
}
