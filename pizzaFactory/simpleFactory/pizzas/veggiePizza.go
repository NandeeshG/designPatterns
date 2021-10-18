package pizzas

import "fmt"

type VeggiePizza struct {
	name   string
	status string
}

func NewVeggiePizza() PizzaInterface {
	return &VeggiePizza{
		name:   "Veggie Pizza",
		status: "Just a order",
	}
}

func (c *VeggiePizza) Prepare() {
	fmt.Println("Preparing the dough for ", c.name)
	c.status = "Dough prepared"
}

func (c *VeggiePizza) Bake() {
	fmt.Println("Baking ", c.name)
	c.status = "Baked"
}

func (c *VeggiePizza) Cut() {
	fmt.Println("Cutting ", c.name)
	c.status = "Prepared and cut"
}

func (c *VeggiePizza) Box() {
	fmt.Println("Boxing ", c.name)
	c.status = "Boxed and packed"
}
