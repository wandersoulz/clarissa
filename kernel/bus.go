package kernel

import (
	"fmt"
	"math"

	"github.com/wandersoulz/godes"
)

// Area - place for Clarissa to interact with people
type Bus struct {
	*godes.Runner
	c           *Clarissa
	destination string
	people      []Friend
	timeOnBus   float64
}

type BusDetails struct {
	PeopleOnBus []Friend
	Destination string
	Clar        *Clarissa
}

func (b *Bus) Run() {
	// Get a random encounter
	encounterProb := randomEncounter.Get(0, 10)
	//fmt.Printf("Clarissa is on the bus on her way to %s\n", b.destination)
	if encounterProb < 3 {
		// we have an encounter
		personIndex := int(peopleIndexDist.Get(0, float64(len(b.people))))
		person := b.people[personIndex]

		friend, isNew := b.c.MakeFriend(&person)
		if isNew {
			fmt.Printf("Clarissa met %s while on the bus\n", person.Name)
		} else {
			fmt.Printf("Clarissa talked to %s on the bus\n", person.Name)
		}
		lengthOfConversation := randomTime.Get(10, 15, b.timeOnBus)
		b.c.GetInfluence(friend, lengthOfConversation/b.timeOnBus)
	}

	godes.Advance(b.timeOnBus)

	if b.destination == "School" {
		// AT SCHOOL, LEARN CLARISSA, LEARN
		godes.AddRunner(GetSchool(b.c))
	} else if b.destination == "Work" {
		// Go to work Clarissa
		godes.AddRunner(GetWork(b.c))
	} else if b.destination == "Home" {
		// Go home Clarissa
		godes.AddRunner(InitHome(b.c))
	}
}

// InitBus - Initialize the bus with a random set of people
func InitBus(destination string, c *Clarissa) *Bus {
	a := Bus{
		&godes.Runner{},
		c,
		destination,
		createPeople(),
		randomTime.Get(30, 33, 40),
	}

	return &a
}

func createPeople() []Friend {
	numPeople := int(math.Floor(numPeopleDist.Get(4, 10, 15)))
	friends := make([]Friend, numPeople)
	for i := 0; i < numPeople; i++ {
		friends[i] = CreateFriend()
	}

	return friends
}
