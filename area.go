package clarissa

import (
	"fmt"
	"math"

	"github.com/wandersoulz/godes"
)

var people_index_dist *godes.UniformDistr
var num_people_dist *godes.TriangularDistr
var random_encounter *godes.UniformDistr
var random_time *godes.TriangularDistr

func InitAreaDists() {
	people_index_dist = godes.NewUniformDistr(true)
	num_people_dist = godes.NewTriangularDistr(true)
	random_encounter = godes.NewUniformDistr(true)
	random_time = godes.NewTriangularDistr(true)
}

// Area - place for Clarissa to interact with people
type Area struct {
	*godes.Runner
	c           *Clarissa
	name        string
	destination string
	people      []Friend
	timeInArea  float64
}

func (a *Area) Run() {
	// Get a random encounter
	encounter := random_encounter.Get(0, 10)
	fmt.Printf("Clarissa is at %s on her way to %s\n", a.name, a.destination)
	if encounter < 5 {
		// we have an encounter
		personIndex := int(people_index_dist.Get(0, float64(len(a.people))))
		person := a.people[personIndex]

		friend, isNew := a.c.MakeFriend(&person)
		if isNew {
			fmt.Printf("Clarissa met %s while on/at the %s\n", person.name, a.name)
		} else {
			fmt.Printf("Clarissa talked to %s on/at the %s\n", person.name, a.name)
		}
		a.c.GetInfluence(friend)
	}
	for godes.GetHour() < 7 && a.destination == "School" {
		godes.Advance(a.timeInArea)
	}
	if godes.GetHour() >= 7 && a.destination == "School" {
		// AT SCHOOL, LEARN CLARISSA, LEARN
		godes.AddRunner(InitSchool(a.c))
	} else if godes.GetHour() >= 13 && a.destination == "Work" {
		// Go to work Clarissa
		godes.AddRunner(InitWork(a.c))
	} else if godes.GetHour() >= 18 && a.destination == "Home" {
		// Go home Clarissa
		godes.AddRunner(InitHome(a.c))
	} else {
		godes.Advance(a.timeInArea)
	}
}

// InitArea - Initialize the area with a random set of people
func InitArea(name, destination string, c *Clarissa) *Area {
	a := Area{
		&godes.Runner{},
		c,
		name,
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
