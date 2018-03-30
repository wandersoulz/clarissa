package kernel

import (
	"fmt"
	"math"

	"github.com/wandersoulz/godes"
)

type School struct {
	*godes.Runner
	c         *Clarissa
	people    []Friend
	schoolEnd float64
}

var school *School

func (s *School) Run() {
	for godes.GetHour() < s.schoolEnd {
		encounterProb := randomEncounter.Get(1, 10)
		if encounterProb > 9 {
			personIndex := int(peopleIndexDist.Get(0, float64(len(s.people))))
			person := s.people[personIndex]

			friend, isNew := s.c.MakeFriend(&person)
			if isNew {
				fmt.Printf("Clarissa met %s while at school\n", person.Name)
			} else {
				fmt.Printf("Clarissa talked to %s at school\n", person.Name)
			}
			maxLength := 40.0
			lengthOfConversation := randomTime.Get(5, 15, maxLength)
			s.c.GetInfluence(friend, lengthOfConversation/maxLength)
			godes.Advance(lengthOfConversation)
		} else {
			lengthOfConversation := randomTime.Get(20, 45, 60)
			godes.Advance(lengthOfConversation)
		}
	}
	godes.AddRunner(InitBus("Work", s.c))
}

func GetSchool(c *Clarissa) *School {
	if school == nil {
		school = &School{
			&godes.Runner{},
			c,
			createChums(),
			14,
		}
	}

	return school
}

func createChums() []Friend {
	numPeople := int(math.Floor(numPeopleDist.Get(15, 17, 30)))
	friends := make([]Friend, numPeople)
	for i := 0; i < numPeople; i++ {
		friends[i] = CreateFriend()
	}

	return friends
}
