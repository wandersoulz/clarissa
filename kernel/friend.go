package kernel

import (
	"github.com/wandersoulz/godes"
	"github.com/wandersoulz/randomname"
)

type Friend struct {
	*godes.PriorityItem
	Preferences []Major
	Name        string
}

func (curr *Friend) Equals(test godes.PriorityInterface) bool {
	testFriend := test.(*Friend)
	return curr.Name == testFriend.Name
}

// CreateFriend creates a new friend with new preferences
func CreateFriend() Friend {
	friend := Friend{
		&godes.PriorityItem{},
		createPreferences(),
		randomname.GetName(),
	}
	return friend
}

func getPreference(dist *godes.UniformDistr) float64 {
	return dist.Get(-30, 30)
}

func createPreferences() []Major {

	cs := NewMajor("Computer Science", getPreference(compsci))
	m := NewMajor("Marketing", getPreference(market))
	a := NewMajor("Astrophysics", getPreference(astro))
	va := NewMajor("Visual Arts", getPreference(visart))
	bio := NewMajor("Biology", getPreference(biology))

	return []Major{cs, m, a, va, bio}
}
