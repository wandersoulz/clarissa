package kernel

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/wandersoulz/godes"
)

// Clarissa - The main protagonist!
type Clarissa struct {
	Preferences []Major
	friends     *godes.PriorityQueue
}

// Log Files
var majorCsvLog *csv.Writer

//var influenceLog *csv.Writer

// Init - Initialize Clarissa!
func Init(seed int64, momFriendship float64) *Clarissa {
	InitKernel()

	majorFilename := fmt.Sprintf("data/major/result_major_%d_%f.csv", seed, momFriendship)
	//influenceFilename := fmt.Sprintf("data/result_influence_%d_%f.csv", seed, momFriendship)

	majorFile, _ := os.Create(majorFilename)
	//influenceFile, _ := os.Create(influenceFilename)
	majorCsvLog = csv.NewWriter(majorFile)
	majorCsvLog.Write([]string{"Time", "Computer Science", "Marketing", "Astrophysics", "Visual Arts", "Biology"})
	//influenceLog = csv.NewWriter(influenceFile)
	//influenceLog.Write([]string{"Time", "Person", "Friendship Score", "Computer Science", "Marketing", "Astrophysics", "Visual Arts", "Biology"})
	friendshipDist = godes.NewTriangularDistr(true)
	scalarDist = godes.NewNormalDistr(true)
	c := Clarissa{
		createPreferences(),
		&godes.PriorityQueue{},
	}

	mom, _ := c.MakeFriend(&Friend{
		&godes.PriorityItem{},
		createPreferences(),
		"Mom",
	})

	c.friends.Update(mom, momFriendship)
	c.GetInfluence(mom, 1)
	return &c
}

// MakeFriend - make a new friend!
func (c Clarissa) MakeFriend(friend *Friend) (*godes.PriorityItem, bool) {
	friendItem := c.friends.Find(friend)
	isNew := false
	if friendItem == nil {
		isNew = true
		friendItem = &godes.PriorityItem{
			Entity:   friend,
			Priority: friendshipDist.Get(-20, 40, 100),
		}
		c.friends.Push(friendItem)
	} else {

		newPriority := friendItem.Priority + friendItem.Priority*scalarDist.Get(.000203, 0.1)
		c.friends.Update(friendItem, newPriority)
	}
	return friendItem, isNew
}

func (c Clarissa) GetInfluence(friend *godes.PriorityItem, interactionMultiplier float64) {
	multiplier := friend.Priority / 10
	friendItem := friend.Entity.(*Friend)

	preferenceString := make([]string, len(c.Preferences)+1)
	preferenceString[0] = fmt.Sprintf("%f", godes.GetSystemTime())
	for j := 0; j < len(c.Preferences); j++ {
		preferenceString[j+1] = strconv.FormatFloat(c.Preferences[j].Value, 'f', 3, 64)
	}
	majorCsvLog.Write(preferenceString)
	majorCsvLog.Flush()

	influenceString := make([]string, len(c.Preferences)+3)
	influenceString[0] = fmt.Sprintf("%f", godes.GetSystemTime())
	influenceString[1] = friend.Entity.(*Friend).Name
	influenceString[2] = strconv.FormatFloat(friend.Priority, 'f', 3, 64)

	for i := 0; i < len(c.Preferences); i++ {
		currentPref := c.Preferences[i].Value
		friendPref := friendItem.Preferences[i].Value
		influenceString[i+3] = strconv.FormatFloat(friendPref, 'f', 3, 64)
		newPref := currentPref + friendPref*multiplier*interactionMultiplier
		c.Preferences[i].Value = newPref
	}

	//influenceLog.Write(influenceString)
	//influenceLog.Flush()

}

func (c Clarissa) ReportCurrentMajor() {
	for i := 0; i < len(c.Preferences); i++ {
		fmt.Printf("%s: %f\n", c.Preferences[i].Name, c.Preferences[i].Value)
	}
}
