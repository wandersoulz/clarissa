package kernel

import (
	"fmt"
	"math"

	"github.com/wandersoulz/godes"
)

var people_index_dist *godes.UniformDistr
var num_people_dist *godes.TriangularDistr
var random_encounter *godes.UniformDistr
var random_time *godes.TriangularDistr

func InitBusDists() {
	people_index_dist = godes.NewUniformDistr(true)
	num_people_dist = godes.NewTriangularDistr(true)
	random_encounter = godes.NewUniformDistr(true)
	random_time = godes.NewTriangularDistr(true)
}

// Area - place for Clarissa to interact with people
type Bus struct {
	*godes.Runner
	c           *Clarissa
	destination string
	people      []Friend
	timeOnBus   float64
}

func (b *Bus) Run() {
	// Get a random encounter
	encounter := random_encounter.Get(0, 10)
	fmt.Printf("Clarissa is on the bus on her way to %s\n", b.destination)
	if encounter < 5 {
		// we have an encounter
		personIndex := int(people_index_dist.Get(0, float64(len(b.people))))
		person := b.people[personIndex]

		friend, isNew := b.c.MakeFriend(&person)
		if isNew {
			fmt.Printf("Clarissa met %s while on the bus\n", person.name)
		} else {
			fmt.Printf("Clarissa talked to %s on the bus\n", person.name)
		}
		b.c.GetInfluence(friend)
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
		random_time.Get(30, 33, 40),
	}

	return &a
}

func createPeople() []Friend {
	numPeople := int(math.Floor(num_people_dist.Get(1, 4, 10)))
	friends := make([]Friend, numPeople)
	for i := 0; i < numPeople; i++ {
		friends[i] = CreateFriend()
	}

	return friends
}
