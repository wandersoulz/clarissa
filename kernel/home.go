package kernel

import (
	"fmt"
	"math"

	"github.com/wandersoulz/godes"
)

type Home struct {
	*godes.Runner
	c         *Clarissa
	sleepTime float64
}

func getFriendshipValues(friends *godes.PriorityQueue) []float64 {
	ret := make([]float64, len(*friends))
	for i := range *friends {
		friend := (*friends)[i]
		ret[i] = friend.Priority
	}
	return ret
}

func getMinMaxFriendValue(friends *godes.PriorityQueue) (float64, float64) {
	min := math.MaxFloat64
	max := float64(math.MinInt64)
	for i := range *friends {
		p := (*friends)[i].Priority
		if p > max {
			max = p
		}
		if p < min {
			min = p
		}
	}
	return min, max

}

func (h *Home) Run() {
	currentDay := godes.GetDay()
	for godes.GetHour() < h.sleepTime && currentDay == godes.GetDay() {
		encounterProb := randomEncounter.Get(1, 10)
		maxLength := 70.0
		lengthOfInteraction := randomTime.Get(30, 40, maxLength)
		if encounterProb > 5 {
			min, max := getMinMaxFriendValue(h.c.friends)
			friendIndex := friendEncDist.Get(getFriendshipValues(h.c.friends), min, max)
			friend := (*h.c.friends)[friendIndex]
			h.c.GetInfluence(friend, lengthOfInteraction/maxLength)
		}

		godes.Advance(lengthOfInteraction)
	}
	fmt.Printf("End of Day: %2.0f\n", currentDay)
	if godes.GetDay() != currentDay {
		godes.AddRunner(InitDay(h.c))
	} else {
		godes.Advance((24 * 60) - (godes.GetHour()*60 + godes.GetMinute()) + 10)
		godes.AddRunner(InitDay(h.c))
	}

}

func InitHome(c *Clarissa) *Home {
	h := &Home{
		&godes.Runner{},
		c,
		sleepTimeDist.Get(20, 23.2, 22.99),
	}

	return h
}
