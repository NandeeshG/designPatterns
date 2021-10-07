package observerSubject

type ObserverInterface interface {
	Update([]int)
	AcceptedBySubject(SubjectInterface)
}
