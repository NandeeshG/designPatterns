package displays

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/weatherStation/observerSubject"
)

type StatisticDisplay struct {
	name string
	sub  observerSubject.SubjectInterface
}

//constructor
func NewStatisticDisplay() *StatisticDisplay {
	return &StatisticDisplay{sub: nil, name: "Forecast"}
}

//displayInterface
func (d *StatisticDisplay) Display() {
	t := d.sub.Get("temp")
	h := d.sub.Get("humi")
	p := d.sub.Get("pres")
	fmt.Printf("Statisitcs: %v.., %v.., %v..\n", t, h, p)
}

//observerInterface
func (d *StatisticDisplay) Update(data []int) {
	fmt.Printf("Statistics recieved: %v\n", data)
	d.Display()
}
func (d *StatisticDisplay) AcceptedBySubject(s observerSubject.SubjectInterface) {
	d.sub = s
}
