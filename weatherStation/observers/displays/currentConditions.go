package displays

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/weatherStation/subjects"
)

type CurrentConditionsDisplay struct {
	sub subjects.SubjectInterface
}

func (d *CurrentConditionsDisplay) Update(data []interface{}) {
	fmt.Printf("Current Conditions recieved: %v\n", data)
	fmt.Printf("d: %v\n", d)
	d.Display()
}

func (d *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current Conditions are: %v\n", d.sub.GetState()...)
}
