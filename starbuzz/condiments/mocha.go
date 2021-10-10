package condiments

import "github.com/NandeeshG/designPatterns/starbuzz/beverages"

type Mocha struct {
	Condiments
}

func NewMocha(b beverages.BeverageInterface) *Mocha {
	ret := &Mocha{}
	rc := &ret.Condiments
	rc.description = "Mocha"
	rc.cost = 5
	rc.beverage = b
	return ret
}
