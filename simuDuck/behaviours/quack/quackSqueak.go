package quack

import "fmt"

type QuackSqueak struct{}

func (q *QuackSqueak) Quack() {
	fmt.Println("Squeak Squeak!")
}
