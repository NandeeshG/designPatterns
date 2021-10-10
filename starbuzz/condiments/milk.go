package condiments

import "github.com/NandeeshG/designPatterns/starbuzz/beverages"

type Milk struct {
	Condiments
}

func NewMilk(b beverages.BeverageInterface) *Milk {
	ret := &Milk{}
	rc := &ret.Condiments
	rc.description = "Milk"
	rc.cost = 35
	rc.beverage = b
	return ret
}
