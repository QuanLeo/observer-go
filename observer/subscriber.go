package observer

type Subscriber interface {
	Start()
	Pause()
	Resume()
	Stop()
}
