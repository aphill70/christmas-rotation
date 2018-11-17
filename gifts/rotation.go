package gifts

import (
	"fmt"
)

type (
	// Rotation represents the gift rotation
	Rotation struct {
		RecipientToGiver map[string]string
		Recipients       map[string]*Gift
		Members          map[string]bool

		Rules *Rules

		currentRecipient *Gift
	}
)

// NewRotation creates a new GiftRotation object
func NewRotation(rulesPath string) (*Rotation, error) {
	rules, err := NewRules(rulesPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to load rules: %v", err)
	}

	return &Rotation{
		Members:          make(map[string]bool),
		Recipients:       make(map[string]*Gift),
		RecipientToGiver: make(map[string]string),

		Rules: rules,
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

	if !r.Members[recipient] {
		r.Members[recipient] = true
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

	if !r.Members[giver] {
		r.Members[giver] = true
	}

	return nil
}

// GetEligibleGivers returns all valid givers for a given recipient
func (r *Rotation) GetEligibleGivers(recipient string) error {
	recipient = normalizeName(recipient)
	if !r.Members[recipient] || r.Recipients[recipient] == nil {
		return fmt.Errorf("Invalid Recipient: %s", recipient)
	}

	gift := r.Recipients[recipient]
	var eligibleMembers = make(map[string]bool)
	for member := range r.Members {
		if member == recipient || gift.Givers[member] || !r.Rules.IsAllowed(recipient, member) {
			continue
		}

		eligibleMembers[member] = true
	}

	fmt.Printf("\n%+v\n", eligibleMembers)

	return nil
}
