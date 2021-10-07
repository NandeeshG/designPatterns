package subjects

import (
	"fmt"
	"log"

	"github.com/NandeeshG/designPatterns/weatherStation/observers"
)

type WeatherData struct {
	temperature int
	humidity    int
	pressure    int
	obs         []observers.ObserverInterface
}

func (w *WeatherData) RegisterObserver(ob observers.ObserverInterface) {
	w.obs = append(w.obs, ob)
	fmt.Printf("w.obs: %v\n", w.obs)
}
func (w *WeatherData) RemoveObserver(ob observers.ObserverInterface) {
	for id, o := range w.obs {
		if o == ob {
			w.obs[id] = w.obs[len(w.obs)-1]
			w.obs = w.obs[:len(w.obs)-1]
		}
	}
}
func (w *WeatherData) NotifyObservers() {
	fmt.Printf("w: %v\n", w)
	for id := range w.obs {
		fmt.Printf("w.obs[id]: %v, %v\n", w.obs[id], &w.obs[id])
		w.obs[id].Update([]interface{}{w.temperature, w.humidity, w.pressure})
	}
}
func (w *WeatherData) GetState() []interface{} {
	return []interface{}{w.temperature, w.humidity, w.pressure}
}
func (w *WeatherData) SetState(data []interface{}) {
	if len(data) < 3 {
		log.Fatalf("Set State must have length at least 3! %v", data...)
		return
	}
	w.temperature = data[0].(int)
	w.humidity = data[1].(int)
	w.pressure = data[2].(int)

	w.NotifyObservers()
}
