package subjects

import "github.com/NandeeshG/designPatterns/weatherStation/observers"

type SubjectInterface interface {
	RegisterObserver(observers.ObserverInterface)
	RemoveObserver(observers.ObserverInterface)
	NotifyObservers()
	GetState() []interface{}
	SetState([]interface{})
}
