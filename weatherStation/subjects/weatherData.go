package subjects

import (
	"fmt"
	"log"

	"github.com/NandeeshG/designPatterns/weatherStation/observerSubject"
)

type WeatherData struct {
	temperature int
	humidity    int
	pressure    int
	obs         []observerSubject.ObserverInterface
}

//constructor
func NewWeatherData() WeatherData {
	return WeatherData{temperature: 0, humidity: 50, pressure: 10}
}

//subjectInterface
func (w *WeatherData) Get(name string) int {
	switch name {
	case "temperature", "temp":
		return w.temperature
	case "humidity", "humi":
		return w.humidity
	case "pressure", "pres":
		return w.pressure
	}
	return 0
}
func (w *WeatherData) GetState() []int {
	return []int{w.temperature, w.humidity, w.pressure}
}
func (w *WeatherData) Set(name string, val int) {
	switch name {
	case "temperature", "temp":
		w.temperature = val
		w.NotifyObservers()
	case "humidity", "humi":
		w.humidity = val
		w.NotifyObservers()
	case "pressure", "pres":
		w.pressure = val
		w.NotifyObservers()
	}
}
func (w *WeatherData) SetState(data []int) {
	if len(data) < 3 {
		log.Fatalf("Set State must have length at least 3! %v", data)
		return
	}
	w.temperature = data[0]
	w.humidity = data[1]
	w.pressure = data[2]

	w.NotifyObservers()
}

func (w *WeatherData) RegisterObserver(ob observerSubject.ObserverInterface) {
	w.obs = append(w.obs, ob)
	ob.AcceptedBySubject(w)
	fmt.Printf("Registered into array. Now w.obs: %v\n", w.obs)
}
func (w *WeatherData) RemoveObserver(ob observerSubject.ObserverInterface) {
	for id := range w.obs {
		if w.obs[id] == ob {
			w.obs[id] = w.obs[len(w.obs)-1]
			w.obs = w.obs[:len(w.obs)-1]
			break
		}
	}
}
func (w *WeatherData) NotifyObservers() {
	for id := range w.obs {
		w.obs[id].Update([]int{w.temperature, w.humidity, w.pressure})
	}
}
