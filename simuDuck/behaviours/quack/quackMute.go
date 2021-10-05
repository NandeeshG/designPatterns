package quack

import "fmt"

type QuackMute struct{}

func (q *QuackMute) Quack() {
	fmt.Println("I am mute :(")
}
