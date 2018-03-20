package kernel

import (
	"fmt"

	"github.com/wandersoulz/godes"
)

var friendEncounterDist *godes.TriangularDistr
var sleepTimeDist *godes.TriangularDistr

func InitHomeDists() {
	friendEncounterDist = godes.NewTriangularDistr(true)
	sleepTimeDist = godes.NewTriangularDistr(true)
}

type Home struct {
	*godes.Runner
	c         *Clarissa
	sleepTime float64
}

func (h *Home) Run() {
	currentDay := godes.GetDay()
	for godes.GetHour() < h.sleepTime && currentDay == godes.GetDay() {
		encounter := random_encounter.Get(1, 10)
		if encounter > 2 {
			friendIndex := int(friendEncounterDist.Get(0, 0, float64(len(*h.c.friends))))
			friend := (*h.c.friends)[friendIndex]
			h.c.GetInfluence(friend)
		}

		lengthOfInteraction := random_time.Get(30, 40, 70)
		godes.Advance(lengthOfInteraction)
	}
	fmt.Printf("End of Day: %2.0f\n", godes.GetDay())
	if godes.GetDay() != currentDay {
		godes.AddRunner(h.c)
	} else {
		godes.Advance((24 * 60) - (godes.GetHour()*60 + godes.GetMinute()))
		godes.AddRunner(h.c)
	}

}

func InitHome(c *Clarissa) *Home {
	h := &Home{
		&godes.Runner{},
		c,
		sleepTimeDist.Get(21, 23.2, 23.99),
	}

	return h
}
