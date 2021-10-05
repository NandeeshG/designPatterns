package ducktypes

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/fly"
	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/quack"
	"github.com/NandeeshG/designPatterns/simuDuck/duck"
)

type DecoyDuck struct {
	duck.Duck
}

func (md *DecoyDuck) Init() {
	md.SetFlyBehaviour(&fly.FlyNoWay{})
	md.SetQuackBehaviour(&quack.QuackMute{})
}

func (md *DecoyDuck) Display() {
	fmt.Println("I am a wooden decoy.")
}
