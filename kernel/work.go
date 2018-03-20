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
		encounter := random_encounter.Get(1, 10)
		if encounter < 4 {
			personIndex := int(people_index_dist.Get(0, float64(len(w.people))))
			person := w.people[personIndex]

			friend, isNew := w.c.MakeFriend(&person)
			if isNew {
				fmt.Printf("Clarissa met %s while at work\n", person.name)
			} else {
				fmt.Printf("Clarissa talked to %s at work\n", person.name)
			}
			w.c.GetInfluence(friend)

			lengthOfConversation := random_time.Get(1, 15, 40)
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
