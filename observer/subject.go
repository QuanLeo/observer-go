package observer

type Subject interface {
	Register(subscriber Subscriber)
	Notification(command Command)
}
