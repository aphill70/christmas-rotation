package gifts

import (
	"fmt"
)

type (
	// Gift represents a recipient and history of givers
	Gift struct {
		Recipient string

		Givers map[string]string
	}
)

// NewGift creates a new Gift to calculate
func NewGift(recipient string) (*Gift, error) {
	if recipient == "" {
		return nil, fmt.Errorf("Recipient must be non empty")
	}

	return &Gift{
		Recipient: recipient,
		Givers:    make(map[string]string),
	}, nil
}

// AddGiver adds a new giver to the Gift
func (g *Gift) AddGiver(giver, year string) {
	g.Givers[year] = giver
}
