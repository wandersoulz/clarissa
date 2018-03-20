package kernel

import (
	"fmt"
)

type EventInterface interface {
	Run()
}

type Event struct {
}

func (e Event) Run() {
	fmt.Println("event is running")
}
