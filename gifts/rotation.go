package gifts

import (
	"fmt"
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
	}, nil
}

// AddRecipient Adds a new person to the rotation
func (r *Rotation) AddRecipient(recipient string) error {
	gift, err := NewGift(recipient)
	if err != nil {
		return fmt.Errorf("Invalid recipient: %s", recipient)
	}
	r.currentRecipient = gift

	if !r.allMembers[recipient] {
		r.allMembers[recipient] = true
	}

	return nil
}

// AddGiver adds a new giver to the current recipient
func (r *Rotation) AddGiver(giver, year string) error {
	if r.currentRecipient == nil {
		return fmt.Errorf("current recipient is null")
	}

	r.currentRecipient.AddGiver(giver, year)

	return nil
}
