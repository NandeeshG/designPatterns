package quack

import "fmt"

type QuackNormal struct{}

func (q *QuackNormal) Quack() {
	fmt.Println("Quack quack!")
}
