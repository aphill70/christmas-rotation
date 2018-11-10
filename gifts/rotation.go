package gifts

import (
	"fmt"
	"strings"
)

type (
	// Rotation represents the gift rotation
	Rotation struct {
		recipientGiver map[string]string

		Recipients map[string]*Gift

		currentRecipient *Gift

		allMembers map[string]bool
	}
)

// NewRotation creates a new GiftRotation object
func NewRotation() (*Rotation, error) {
	return &Rotation{
		allMembers: make(map[string]bool),
		Recipients: make(map[string]*Gift),
	}, nil
}

// AddRecipient Adds a new person to the rotation
func (r *Rotation) AddRecipient(recipient string) error {
	recipient = normalizeName(recipient)
	gift, err := NewGift(recipient)
	if err != nil {
		return fmt.Errorf("Invalid recipient: %s", recipient)
	}

	r.currentRecipient = gift
	r.Recipients[recipient] = gift

	if !r.allMembers[recipient] {
		r.allMembers[recipient] = true
	}

	return nil
}

// AddGiver adds a new giver to the current recipient
func (r *Rotation) AddGiver(giver, year string) error {
	giver = normalizeName(giver)
	if r.currentRecipient == nil {
		return fmt.Errorf("current recipient is null")
	}

	r.currentRecipient.AddGiver(giver, year)

	if !r.allMembers[giver] {
		r.allMembers[giver] = true
	}

	return nil
}

// GetEligibleGivers returns all valid givers for a given recipient
func (r *Rotation) GetEligibleGivers(recipient string) error {
	recipient = normalizeName(recipient)
	if !r.allMembers[recipient] || r.Recipients[recipient] == nil {
		return fmt.Errorf("Invalid Recipient: %s", recipient)
	}

	gift := r.Recipients[recipient]
	var eligibleMembers = make(map[string]bool)
	for member := range r.allMembers {
		if member == recipient || gift.altGivers[member] {
			continue
		}

		eligibleMembers[member] = true
	}

	fmt.Printf("\n%+v\n", eligibleMembers)

	return nil
}

func normalizeName(name string) string {
	normalized := strings.Trim(strings.ToLower(name), " ")

	if strings.HasPrefix(normalized, "chris") {
		return "christopher"
	}

	return normalized
}
