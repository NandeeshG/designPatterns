package main

import (
	"fmt"

	"github.com/NandeeshG/designPatterns/weatherStation/displays"
	"github.com/NandeeshG/designPatterns/weatherStation/observerSubject"
	"github.com/NandeeshG/designPatterns/weatherStation/subjects"
)

func main() {
	wd := subjects.WeatherData{}
	fmt.Printf("MAIN wd: %v\n", wd)

	obs := []observerSubject.ObserverInterface{
		displays.NewCurrentConditionsDisplay(),
		displays.NewStatisticDisplay(),
		displays.NewForecastDisplay(),
	}
	fmt.Printf("MAIN obs: %v\n", obs)

	for id, _ := range obs {
		wd.RegisterObserver(obs[id])
	}

	wd.SetState([]int{10, 20, 30})
	wd.SetState([]int{100, 200, 300})
	wd.RemoveObserver(obs[0])
	wd.SetState([]int{1000, 2000, 3000})
}
