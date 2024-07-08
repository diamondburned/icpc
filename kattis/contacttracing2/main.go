package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type InfectionStatus uint8

const (
	Uninfected InfectionStatus = iota
	InfectedNotContagious
	InfectedContagious
	InfectedQuarantined
)

type Person struct {
	ID     int
	Status InfectionStatus
}

// People is a list of people.
type People []Person

// WithID returns the person with the given ID.
func (p People) WithID(id int) *Person {
	return &p[id-1]
}

type Day struct {
	// Contacts is a map of who contacted whom. Beware! This is a bidirectional
	// map, so if A contacts B, then B contacts A!
	Contacts Contacts
}

// Contacts is a map of who contacted whom. Beware! This is a bidirectional
// map, so if A contacts B, then B contacts A!
type Contacts map[int]map[int]bool

// EachContacted calls the given function for each person who was contacted on
// this day.
func (c Contacts) EachContacted(f func(id int)) {
	contacted := make(map[int]bool, len(c))
	for a, whoms := range c {
		if !contacted[a] {
			contacted[a] = true
			f(a)
		}

		for b := range whoms {
			if !contacted[b] {
				contacted[b] = true
				f(b)
			}
		}
	}
}

// Add adds a new pair of contacts.
func (c *Contacts) Add(a, b int) {
	if *c == nil {
		*c = make(map[int]map[int]bool, 2)
	}

	aContacts, ok := (*c)[a]
	if !ok {
		aContacts = make(map[int]bool, 1)
		(*c)[a] = aContacts
	}
	aContacts[b] = true

	bContacts, ok := (*c)[b]
	if !ok {
		bContacts = make(map[int]bool, 1)
		(*c)[b] = bContacts
	}
	bContacts[a] = true
}

// HasContacted returns true if the person was contacted by someone else.
func (c Contacts) HasContacted(a, b int) bool {
	return c[a][b] || c[b][a]
}

func main() {
	var npeople, today, ncontacts int
	fmt.Scan(&npeople, &today, &ncontacts)

	people := make(People, npeople)
	for i := range people {
		people[i] = Person{ID: i + 1, Status: Uninfected}
	}

	// Assume person 1 is infected. They may not necessarily be patient zero,
	// but they'll be spreading the disease.
	people.WithID(1).Status = InfectedContagious

	days := make([]Day, today)

	for i := 0; i < ncontacts; i++ {
		var who, whom, when int
		fmt.Scan(&who, &whom, &when)

		day := &days[when-1]
		day.Contacts.Add(who, whom)
	}

	spew.Dump(days)
	simulate(people, days)
}

func simulate(people People, days []Day) {
	for _, day := range days {
		// First, find everyone who the person might've infected. These people
		// might be infected, but they're not contagious yet.
		var possiblyInfected []int
		for a, to := range day.Contacts {
			if people.WithID(a).Status != InfectedContagious {
				continue
			}
			for b := range to {
				possiblyInfected = append(possiblyInfected, b)
			}
		}

		// Simulate every possible infection: if person A is actually infected,
		// then who they contacted is infected.
		for _, id := range possiblyInfected {
			// Make a copy of everyone. Assume this is a different universe
			// where this one person is infected.
			clone := append(People(nil), people...)
			clone.WithID(id).Status = InfectedNotContagious
		}

		for a, b := range day.Contacts {
			if people.WithID(a).Status == InfectedContagious {
			}
		}
	}
}
