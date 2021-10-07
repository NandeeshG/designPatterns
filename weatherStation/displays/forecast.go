package displays

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/weatherStation/observerSubject"
)

type ForecastDisplay struct {
	name string
	sub  observerSubject.SubjectInterface
}

//constructor
func NewForecastDisplay() *ForecastDisplay {
	return &ForecastDisplay{sub: nil, name: "Forecast"}
}

//displayInterface
func (d *ForecastDisplay) Display() {
	t := d.sub.Get("temp")
	h := d.sub.Get("humi")
	p := d.sub.Get("pres")
	fmt.Printf("Forecasting: %v, %v, %v\n", t+1, h+2, p+3)
}

//observerInterface
func (d *ForecastDisplay) Update(data []int) {
	fmt.Printf("Forecast recieved: %v\n", data)
	d.Display()
}
func (d *ForecastDisplay) AcceptedBySubject(s observerSubject.SubjectInterface) {
	d.sub = s
}
