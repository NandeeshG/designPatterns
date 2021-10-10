package beverages

type DarkRoast struct {
	description string
	cost        int
}

func NewDarkRoast() *DarkRoast {
	ret := DarkRoast{
		description: "DarkRoast Drink",
		cost:        30,
	}
	return &ret
}

func (d *DarkRoast) GetDescription() string {
	return d.description
}

func (d *DarkRoast) GetCost() int {
	return d.cost
}
