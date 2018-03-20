package kernel

import (
	"github.com/icrowley/fake"
	"github.com/wandersoulz/godes"
)

type Friend struct {
	*godes.PriorityItem
	preferences []Major
	name        string
}

func (curr *Friend) Equals(test godes.PriorityInterface) bool {
	testFriend := test.(*Friend)
	return curr.name == testFriend.name
}

// CreateFriend creates a new friend with new preferences
func CreateFriend() Friend {
	friend := Friend{
		&godes.PriorityItem{},
		createPreferences(),
		fake.FirstName(),
	}
	return friend
}

func createPreferences() []Major {
	cs := NewMajor("Computer Science", comp_sci.Get(-20, 20))
	m := NewMajor("Marketing", market.Get(-20, 20))
	a := NewMajor("Astrophysics", astro.Get(-20, 20))
	va := NewMajor("Visual Arts", vis_art.Get(-20, 20))
	bio := NewMajor("Biology", biology.Get(-20, 20))

	return []Major{cs, m, a, va, bio}
}
