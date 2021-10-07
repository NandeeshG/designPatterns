package observerSubject

type SubjectInterface interface {
	RegisterObserver(ObserverInterface)
	RemoveObserver(ObserverInterface)
	NotifyObservers()

	Get(string) int
	GetState() []int

	Set(string, int)
	SetState([]int)
}
