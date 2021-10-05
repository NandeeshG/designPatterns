package fly

import "fmt"

type FlyWithWings struct {
}

func (f *FlyWithWings) Fly() {
	fmt.Println("Flying with wings!")
}
