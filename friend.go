package clarissa

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
	cs := NewMajor("Computer Science", comp_sci.Get(-10, 10))
	m := NewMajor("Marketing", market.Get(-10, 10))
	a := NewMajor("Astrophysics", astro.Get(-10, 10))
	va := NewMajor("Visual Arts", vis_art.Get(-10, 10))
	bio := NewMajor("Biology", biology.Get(-10, 10))

	return []Major{cs, m, a, va, bio}
}
