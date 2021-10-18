package simplepizza

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/pizzaFactory/simpleFactory/pizzas"
)

type PizzaStore struct {
	factory *SimplePizzaFactory
}

func NewPizzaStore(spf *SimplePizzaFactory) *PizzaStore {
	pizzaStore := PizzaStore{factory: spf}
	return &pizzaStore
}

func (ps *PizzaStore) OrderPizza(pizzaType string) pizzas.PizzaInterface {
	pizza := ps.factory.CreatePizza(pizzaType)
	if pizza == nil {
		fmt.Println("Unable to create pizza")
		return nil
	}
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
