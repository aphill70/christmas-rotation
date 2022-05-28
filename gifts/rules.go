package gifts

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	//Rules defines the rotation rules
	Rules struct {
		Households [][]string `json:"households"`

		PersonLookup map[string]map[string]bool
	}
)

// NewRules reads a file from path and returns the rule object
func NewRules(path string) (*Rules, error) {
	rules := &Rules{
		PersonLookup: make(map[string]map[string]bool),
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read rules file: %v", err)
	}

	err = json.NewDecoder(f).Decode(&rules)
	if err != nil {
		return nil, fmt.Errorf("failed to decode rules file: %v", err)
	}

	for _, household := range rules.Households {
		for _, person := range household {
			rules.PersonLookup[person] = make(map[string]bool)

			for _, member := range household {
				if member != person {
					rules.PersonLookup[person][member] = true
				}
			}
		}
	}

	return rules, nil
}

// IsAllowed returns true if the rules do not forbid them from recieving a gift from the gifter
func (r *Rules) IsAllowed(recipient, gifter string) bool {
	if r == nil || r.PersonLookup == nil || r.PersonLookup[recipient] == nil {
		return true
	}
	return !r.PersonLookup[recipient][gifter]
}
