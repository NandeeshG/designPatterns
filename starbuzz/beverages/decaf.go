package beverages

type Decaf struct {
	description string
	cost        int
}

func NewDecaf() *Decaf {
	ret := Decaf{
		description: "Decaf Drink",
		cost:        10,
	}
	return &ret
}

func (d *Decaf) GetDescription() string {
	return d.description
}

func (d *Decaf) GetCost() int {
	return d.cost
}
