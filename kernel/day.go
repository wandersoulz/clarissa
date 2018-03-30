package kernel

import "github.com/wandersoulz/godes"

type NewDay struct {
	*godes.Runner
	c *Clarissa
}

func (d *NewDay) Run() {
	simTime := godes.GetSimulationTime()
	if godes.GetDay() >= 366 {
		return
	}
	currentTime := simTime.Hour*60 + simTime.Minute
	godes.Advance(6.5*60 - currentTime)
	// Clarissa is awake and is at home and ready to go
	// She gets on the bus
	godes.AddRunner(InitBus("School", d.c))
}

// InitDay - start a new day
func InitDay(c *Clarissa) *NewDay {
	return &NewDay{
		&godes.Runner{},
		c,
	}
}
