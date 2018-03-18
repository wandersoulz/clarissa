package clarissa

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
	godes.AddRunner(InitArea("Bus", "Home", w.c))
}

func InitWork(c *Clarissa) *Work {
	w := &Work{
		&godes.Runner{},
		c,
		createChums(),
		18,
	}
	return w
}
