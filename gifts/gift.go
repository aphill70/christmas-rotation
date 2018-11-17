package gifts

import (
	"fmt"
)

type (
	// Gift represents a recipient and history of givers
	Gift struct {
		Recipient string

		History map[string]string
		Givers  map[string]bool
	}
)

// NewGift creates a new Gift to calculate
func NewGift(recipient string) (*Gift, error) {
	if recipient == "" {
		return nil, fmt.Errorf("Recipient must be non empty")
	}

	return &Gift{
		Recipient: normalizeName(recipient),
		History:   make(map[string]string),
		Givers:    make(map[string]bool),
	}, nil
}

// AddGiver adds a new giver to the Gift
func (g *Gift) AddGiver(giver, year string) error {
	normalizedGiver := normalizeName(giver)
	if g.History[year] != "" {
		return fmt.Errorf("cannot duplicate years. giver for year %s already exists", year)
	}
	g.History[year] = normalizedGiver

	if g.Givers[normalizedGiver] {
		g.Givers = map[string]bool{
			normalizedGiver: true,
		}
	} else {
		g.Givers[normalizedGiver] = true
	}

	return nil
}
