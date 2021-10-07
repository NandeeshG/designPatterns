package displays

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/weatherStation/subjects"
)

type StatisticsDisplay struct {
	sub subjects.SubjectInterface
}

func (d *StatisticsDisplay) Update(data []interface{}) {
	fmt.Printf("Statistics recieved: %v\n", data)
}

func (d *StatisticsDisplay) Display() {
	fmt.Printf("Statistics are: %v\n", d.sub.GetState()...)
}
