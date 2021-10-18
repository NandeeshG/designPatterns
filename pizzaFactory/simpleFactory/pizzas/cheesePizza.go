package pizzas

import "fmt"

type CheesePizza struct {
	name   string
	status string
}

func NewCheesePizza() PizzaInterface {
	return &CheesePizza{
		name:   "Cheese Pizza",
		status: "Just a order",
	}
}

func (c *CheesePizza) Prepare() {
	fmt.Println("Preparing the dough for ", c.name)
	c.status = "Dough prepared"
}

func (c *CheesePizza) Bake() {
	fmt.Println("Baking ", c.name)
	c.status = "Baked"
}

func (c *CheesePizza) Cut() {
	fmt.Println("Cutting ", c.name)
	c.status = "Prepared and cut"
}

func (c *CheesePizza) Box() {
	fmt.Println("Boxing ", c.name)
	c.status = "Boxed and packed"
}
