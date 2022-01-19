package observer

type Command int

const (
	START  Command = 0
	PAUSE          = 1
	RESUME         = 2
	STOP           = 3
)
