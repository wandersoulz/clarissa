package clarissa

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/wandersoulz/godes"
)

// Clarissa - The main protagonist!
type Clarissa struct {
	*godes.Runner
	preferences []Major
	friends     *godes.PriorityQueue
	isAwake     bool
}

// Log Files
var majorCsvLog *csv.Writer
var influenceLog *csv.Writer

var friendshipDist *godes.TriangularDistr
var scalarDist *godes.NormalDistr

// Run - Clarissa's run method
func (c *Clarissa) Run() {
	simTime := godes.GetSimulationTime()
	if godes.GetDay() == 366 {
		return
	}
	if simTime.Hour < 6 {
		currentTime := simTime.Hour*60 + simTime.Minute
		godes.Advance(6*60 - currentTime)
	}
	if simTime.Hour >= 6 && !c.isAwake {
		// Time to wake up!
		c.isAwake = true
		godes.Advance(30 - simTime.Minute)
	}
	// Clarissa is awake and is at home and ready to go
	// She gets on the bus
	godes.AddRunner(InitArea("Bus", "School", c))
}

// Init - Initialize Clarissa!
func Init() *Clarissa {
	majorFile, _ := os.Create("result_major.csv")
	influenceFile, _ := os.Create("result.csv")
	majorCsvLog = csv.NewWriter(majorFile)
	majorCsvLog.Write([]string{"Time", "Computer Science", "Marketing", "Astrophysics", "Visual Arts", "Biology"})
	influenceLog = csv.NewWriter(influenceFile)
	influenceLog.Write([]string{"Time", "Person", "Friendship Score", "Computer Science", "Marketing", "Astrophysics", "Visual Arts", "Biology"})
	friendshipDist = godes.NewTriangularDistr(true)
	scalarDist = godes.NewNormalDistr(true)
	c := Clarissa{
		&godes.Runner{},
		createPreferences(),
		&godes.PriorityQueue{},
		false,
	}

	mom, _ := c.MakeFriend(&Friend{
		&godes.PriorityItem{},
		createPreferences(),
		"Mom",
	})
	momPriority, _ := strconv.ParseFloat(os.Args[2], 64)
	c.friends.Update(mom, momPriority)
	c.GetInfluence(mom)
	return &c
}

// MakeFriend - ake a new friend!
func (c Clarissa) MakeFriend(friend *Friend) (*godes.PriorityItem, bool) {
	friendItem := c.friends.Find(friend)
	isNew := false
	if friendItem == nil {
		isNew = true
		friendItem = &godes.PriorityItem{
			Entity:   friend,
			Priority: friendshipDist.Get(-2, 4, 10),
		}
		c.friends.Push(friendItem)
	} else {
		newPriority := friendItem.Priority * scalarDist.Get(1.0003, 0.01)
		c.friends.Update(friendItem, newPriority)
	}
	return friendItem, isNew
}

func (c Clarissa) GetInfluence(friend *godes.PriorityItem) {
	multiplier := friend.Priority
	friendItem := friend.Entity.(*Friend)

	fudgeFactor := scalarDist.Get(0.6, 0.01)

	preferenceString := make([]string, len(c.preferences)+1)
	preferenceString[0] = fmt.Sprintf("%f", godes.GetSystemTime())
	for j := 0; j < len(c.preferences); j++ {
		preferenceString[j+1] = strconv.FormatFloat(c.preferences[j].value, 'f', 3, 64)
	}
	majorCsvLog.Write(preferenceString)
	majorCsvLog.Flush()

	influenceString := make([]string, len(c.preferences)+3)
	influenceString[0] = fmt.Sprintf("%f", godes.GetSystemTime())
	influenceString[1] = friend.Entity.(*Friend).name
	influenceString[2] = strconv.FormatFloat(friend.Priority, 'f', 3, 64)

	for i := 0; i < len(c.preferences); i++ {
		currentPref := c.preferences[i].value
		friendPref := friendItem.preferences[i].value
		influenceString[i+3] = strconv.FormatFloat(friendPref, 'f', 3, 64)
		newPref := currentPref + friendPref*multiplier*fudgeFactor
		c.preferences[i].value = newPref
	}

	influenceLog.Write(influenceString)
	influenceLog.Flush()

}

func (c Clarissa) ReportCurrentMajor() {
	for i := 0; i < len(c.preferences); i++ {
		fmt.Printf("%s: %f\n", c.preferences[i].name, c.preferences[i].value)
	}
}
