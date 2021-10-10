package condiments

import "github.com/NandeeshG/designPatterns/starbuzz/beverages"

type Soy struct {
	Condiments
}

func NewSoy(b beverages.BeverageInterface) *Soy {
	ret := &Soy{}
	rc := &ret.Condiments
	rc.description = "Soy"
	rc.cost = 25
	rc.beverage = b
	return ret
}
