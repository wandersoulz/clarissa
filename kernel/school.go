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
		encounter := random_encounter.Get(1, 10)
		if encounter > 4 {
			personIndex := int(people_index_dist.Get(0, float64(len(s.people))))
			person := s.people[personIndex]

			friend, isNew := s.c.MakeFriend(&person)
			if isNew {
				fmt.Printf("Clarissa met %s while at school\n", person.name)
			} else {
				fmt.Printf("Clarissa talked to %s at school\n", person.name)
			}
			s.c.GetInfluence(friend)

			lengthOfConversation := random_time.Get(10, 15, 40)
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
	numPeople := int(math.Floor(num_people_dist.Get(15, 17, 30)))
	friends := make([]Friend, numPeople)
	for i := 0; i < numPeople; i++ {
		friends[i] = CreateFriend()
	}

	return friends
}
