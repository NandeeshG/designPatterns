package main

import simplepizza "github.com/NandeeshG/designPatterns/pizzaFactory/simpleFactory/simplePizza"

func main() {
	pizzaFactory := simplepizza.NewSimplePizzaFactory()
	pizzaStore := simplepizza.NewPizzaStore(pizzaFactory)

	simulatingOrders := []string{
		"veggiePizza",
		"cheesePizza",
		"cheesePizza",
		"veggiePizza",
	}

	for _, order := range simulatingOrders {
		pizzaStore.OrderPizza(order)
	}
}
