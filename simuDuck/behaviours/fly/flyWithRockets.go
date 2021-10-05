package fly

import "fmt"

type FlyWithRockets struct{}

func (f *FlyWithRockets) Fly() {
	fmt.Println("Rocketman!")
}
