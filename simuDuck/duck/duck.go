package duck

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/fly"
	"github.com/NandeeshG/designPatterns/simuDuck/behaviours/quack"
)

type Duck struct {
	FlyBehaviour   fly.FlyBehaviour
	QuackBehaviour quack.QuackBehaviour

	Display func()
}

func (d *Duck) PerformFly() {
	d.FlyBehaviour.Fly()
}

func (d *Duck) PerformQuack() {
	d.QuackBehaviour.Quack()
}

func (d *Duck) Swim() {
	fmt.Println("Every duck swims :)")
}
