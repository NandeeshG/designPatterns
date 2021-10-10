package beverages

type Espresso struct {
	description string
	cost        int
}

func NewEspresso() *Espresso {
	ret := Espresso{
		description: "Espresso Drink",
		cost:        20,
	}
	return &ret
}

func (d *Espresso) GetDescription() string {
	return d.description
}

func (d *Espresso) GetCost() int {
	return d.cost
}
