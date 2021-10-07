package main

import (
	"github.com/NandeeshG/designPatterns/weatherStation/observers"
	"github.com/NandeeshG/designPatterns/weatherStation/observers/displays"
	"github.com/NandeeshG/designPatterns/weatherStation/subjects"
)

func main() {
	wd := subjects.WeatherData{}

	obs := []observers.ObserverInterface{
		&displays.CurrentConditionsDisplay{},
		&displays.StatisticsDisplay{},
		&displays.ForecastDisplay{},
	}

	for id, _ := range obs {
		wd.RegisterObserver(obs[id])
	}

	wd.SetState([]interface{}{10, 20, 30})
	wd.SetState([]interface{}{100, 200, 300})
	wd.RemoveObserver(obs[0])
	wd.SetState([]interface{}{1000, 2000, 3000})
}
