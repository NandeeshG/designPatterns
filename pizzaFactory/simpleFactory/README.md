## Simple Pizza Factory

1. We delegate pizza creation job to SimplePizzaFactory
2. Pizzastore just calls the factory, so it doesn't need to worry about that.
3. Changes that can happen:
    * Add more stores like DelhiPizzaStore, MumbaiPizzaStore
    * Add more PizzaFactories that handle more varied types of Pizzas
* To handle such changes we can encapsulate pizzaStores and Factories again in their own interface and then write code to those interfaces.
