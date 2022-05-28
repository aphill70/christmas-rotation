package gifts

import (
	"fmt"
)

type (
	// Gift represents a recipient and history of givers
	Gift struct {
		Recipient string

		History     map[string]*Giver
		Givers      []*Giver
		GiverLookup map[string]*Giver
	}

	// Giver struct for representing a giver and the number of times they have given
	Giver struct {
		Count int
		Giver string
	}
)

// GetGiver returns the giver info entry
func (g *Gift) GetGiver(giver string) *Giver {
	for _, entry := range g.Givers {
		if giver == entry.Giver {
			return entry
		}
	}

	return nil
}

// NewGift creates a new Gift to calculate
func NewGift(recipient string) (*Gift, error) {
	if recipient == "" {
		return nil, fmt.Errorf("recipient must be non empty")
	}

	return &Gift{
		Recipient:   NormalizeName(recipient),
		History:     make(map[string]*Giver),
		Givers:      []*Giver{},
		GiverLookup: make(map[string]*Giver),
	}, nil
}

// AddGiver adds a new giver to the Gift
func (g *Gift) AddGiver(giver, year string) error {
	normalizedGiver := NormalizeName(giver)
	if g.History[year] != nil {
		return fmt.Errorf("cannot duplicate years. giver for year %s already exists", year)
	}

	giverStruct := g.GiverLookup[normalizedGiver]
	if giverStruct == nil {
		giverStruct = &Giver{
			Count: 1,
			Giver: normalizedGiver,
		}
		g.GiverLookup[normalizedGiver] = giverStruct
		g.Givers = append(g.Givers, giverStruct)
	} else {
		giverStruct.Count++
	}

	g.History[year] = giverStruct

	return nil
}

// GetCurrentRotationIndex d
func (g *Gift) GetCurrentRotationIndex() int {
	currentRotationIndex := 0
	init := false
	for _, val := range g.GiverLookup {
		if !init {
			currentRotationIndex = val.Count
			init = true
		} else if val.Count < currentRotationIndex {
			currentRotationIndex = val.Count
		}
	}
	return currentRotationIndex
}
