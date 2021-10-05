package fly

import "fmt"

type FlyNoWay struct {
}

func (f *FlyNoWay) Fly() {
	fmt.Println("I cannot fly...")
}
