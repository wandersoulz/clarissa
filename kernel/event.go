package kernel

import (
	"fmt"
)

type EventType string

const (
	EncounterEvent EventType = "encouter"
	BusEvent       EventType = "bus"
	SchoolEvent    EventType = "school"
	WorkEvent      EventType = "work"
	HomeEvent      EventType = "home"
	NewDayEvent    EventType = "newday"
)

// EventInterface - event interface
type EventInterface interface {
	Run()
}

// Event - exported event type
type Event struct {
	Type    EventType
	Time    float64
	Details interface{}
}

func (e Event) Run() {
	fmt.Println("event is running")
}
