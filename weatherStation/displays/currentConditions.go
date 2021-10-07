package displays

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/weatherStation/observerSubject"
)

type CurrentConditionsDisplay struct {
	name string
	sub  observerSubject.SubjectInterface
}

//constructor
func NewCurrentConditionsDisplay() *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{sub: nil, name: "Current"}
}

//displayInterface
func (d *CurrentConditionsDisplay) Display() {
	t := d.sub.Get("temp")
	h := d.sub.Get("humi")
	p := d.sub.Get("pres")
	fmt.Printf("Current conditions: %v, %v, %v\n", t, h, p)
}

//observerInterface
func (d *CurrentConditionsDisplay) Update(data []int) {
	fmt.Printf("Current Conditions recieved: %v\n", data)
	d.Display()
}
func (d *CurrentConditionsDisplay) AcceptedBySubject(s observerSubject.SubjectInterface) {
	d.sub = s
}
