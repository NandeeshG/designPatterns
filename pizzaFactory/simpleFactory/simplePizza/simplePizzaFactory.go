package simplepizza

import "github.com/NandeeshG/designPatterns/pizzaFactory/simpleFactory/pizzas"

type SimplePizzaFactory struct {
}

func NewSimplePizzaFactory() *SimplePizzaFactory {
	return &SimplePizzaFactory{}
}

func (spf *SimplePizzaFactory) CreatePizza(pizzaType string) pizzas.PizzaInterface {
	if pizzaType == "veggiePizza" {
		return pizzas.NewVeggiePizza()
	} else if pizzaType == "cheesePizza" {
		return pizzas.NewCheesePizza()
	} else {
		return nil
	}
}
