package beverages

type HouseBlend struct {
	description string
	cost        int
}

func NewHouseBlend() *HouseBlend {
	ret := HouseBlend{
		description: "HouseBlend Drink",
		cost:        40,
	}
	return &ret
}

func (d *HouseBlend) GetDescription() string {
	return d.description
}

func (d *HouseBlend) GetCost() int {
	return d.cost
}
