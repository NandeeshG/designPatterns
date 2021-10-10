package condiments

import "github.com/NandeeshG/designPatterns/starbuzz/beverages"

type Whip struct {
	Condiments
}

func NewWhip(b beverages.BeverageInterface) *Whip {
	ret := Whip{}
	rc := &ret.Condiments
	rc.description = "Whip"
	rc.cost = 15
	rc.beverage = b
	return &ret
}
