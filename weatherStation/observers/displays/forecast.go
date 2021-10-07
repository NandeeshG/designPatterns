package displays

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/weatherStation/subjects"
)

type ForecastDisplay struct {
	sub subjects.SubjectInterface
}

func (d *ForecastDisplay) Update(data []interface{}) {
	fmt.Printf("Forecast recieved: %v\n", data)
}

func (d *ForecastDisplay) Display() {
	fmt.Printf("Forecast are: %v\n", d.sub.GetState()...)
}
