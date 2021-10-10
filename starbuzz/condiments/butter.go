package condiments

import (
	"github.com/NandeeshG/designPatterns/starbuzz/beverages"
)

//Emulates inheritance
type Butter struct {
	Condiments
}

func NewButter(b beverages.BeverageInterface) *Butter {
	ret := Butter{}
	rc := &ret.Condiments
	rc.description = "Butter"
	rc.cost = 2
	rc.beverage = b
	return &ret
}
