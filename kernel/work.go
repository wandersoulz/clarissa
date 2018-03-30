package kernel

import (
	"fmt"

	"github.com/wandersoulz/godes"
)

type Work struct {
	*godes.Runner
	c       *Clarissa
	people  []Friend
	workEnd float64
}

var work *Work

func (w *Work) Run() {
	for godes.GetHour() < w.workEnd {
		encounterProb := randomEncounter.Get(1, 10)
		if encounterProb < 6 {
			personIndex := int(peopleIndexDist.Get(0, float64(len(w.people))))
			person := w.people[personIndex]

			friend, isNew := w.c.MakeFriend(&person)
			if isNew {
				fmt.Printf("Clarissa met %s while at work\n", person.Name)
			} else {
				fmt.Printf("Clarissa talked to %s at work\n", person.Name)
			}
			maxLength := 40.0
			lengthOfConversation := randomTime.Get(5, 15, maxLength)
			w.c.GetInfluence(friend, lengthOfConversation/maxLength)

			godes.Advance(lengthOfConversation)

		} else {
			maxLength := 100.0
			lengthOfConversation := randomTime.Get(20, 35, maxLength)
			godes.Advance(lengthOfConversation)
		}
	}
	godes.AddRunner(InitBus("Home", w.c))
}

func GetWork(c *Clarissa) *Work {
	if work == nil {
		work = &Work{
			&godes.Runner{},
			c,
			createChums(),
			18,
		}
	}
	return work
}
