package observer

import (
	"fmt"
)

type SubjectMonitor struct {
	subscribers []Subscriber
}

func (subjectMonitor *SubjectMonitor) Register(subscriber Subscriber) {
	subjectMonitor.subscribers = append(subjectMonitor.subscribers, subscriber)
}

func (subjectMonitor *SubjectMonitor) Notification(command Command) {
	switch command {
	case START:
		subjectMonitor.handleStart()
	case PAUSE:
		subjectMonitor.handlePause()
	case RESUME:
		subjectMonitor.handleResume()
	case STOP:
		subjectMonitor.handleStop()
	default:
		fmt.Println("Command not found. Command : ", command)
	}
}

func (subjectMonitor *SubjectMonitor) handleStart() {
	for _, subscriber := range subjectMonitor.subscribers {
		subscriber.Start()
	}
}

func (subjectMonitor *SubjectMonitor) handlePause() {
	for _, subscriber := range subjectMonitor.subscribers {
		subscriber.Pause()
	}
}

func (subjectMonitor *SubjectMonitor) handleResume() {
	for _, subscriber := range subjectMonitor.subscribers {
		subscriber.Resume()
	}
}

func (subjectMonitor *SubjectMonitor) handleStop() {
	for _, subscriber := range subjectMonitor.subscribers {
		subscriber.Stop()
	}
}
