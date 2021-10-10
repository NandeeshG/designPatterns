package main

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/starbuzz/beverages"
	"github.com/NandeeshG/designPatterns/starbuzz/condiments"
)

func main() {
	//customer0 buys
	// butter butter darkRoast
	var drink0 beverages.BeverageInterface = beverages.NewDarkRoast()
	drink0 = condiments.NewButter(drink0)
	drink0 = condiments.NewButter(drink0)
	fmt.Println(drink0.GetDescription())
	fmt.Println(drink0.GetCost())

	//customer 1 buys
	// double mocha houseBlend
	var drink1 beverages.BeverageInterface = beverages.NewHouseBlend()
	drink1 = condiments.NewMocha(drink1)
	drink1 = condiments.NewMocha(drink1)
	fmt.Println(drink1.GetDescription())
	fmt.Println(drink1.GetCost())

	//customer 2 buys
	// whip milk espresso
	var drink2 beverages.BeverageInterface = beverages.NewEspresso()
	drink2 = condiments.NewMilk(drink2)
	drink2 = condiments.NewWhip(drink2)
	fmt.Println(drink2.GetDescription())
	fmt.Println(drink2.GetCost())

	//customer 3 buys
	// triple soy double mocha triple milk decaf
	var drink3 beverages.BeverageInterface = beverages.NewDecaf()
	drink3 = condiments.NewSoy(drink3)
	drink3 = condiments.NewSoy(drink3)
	drink3 = condiments.NewSoy(drink3)
	drink3 = condiments.NewMocha(drink3)
	drink3 = condiments.NewMocha(drink3)
	drink3 = condiments.NewMilk(drink3)
	drink3 = condiments.NewMilk(drink3)
	drink3 = condiments.NewMilk(drink3)
	fmt.Println(drink3.GetDescription())
	fmt.Println(drink3.GetCost())
}
