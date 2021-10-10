package condiments

import "github.com/NandeeshG/designPatterns/starbuzz/beverages"

type Condiments struct {
	beverage    beverages.BeverageInterface
	description string
	cost        int
}

func NewCondiments(str string, c int, b beverages.BeverageInterface) *Condiments {
	ret := Condiments{description: str, cost: c, beverage: b}
	return &ret
}

func (c *Condiments) GetDescription() string {
	return c.description + " " + c.beverage.GetDescription()
}

func (c *Condiments) GetCost() int {
	return c.cost + c.beverage.GetCost()
}
