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
	}
)

// NewRules reads a file from path and returns the rule object
func NewRules(path string) (*Rules, error) {
	rules := &Rules{}

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read rules file: %v", err)
	}

	err = json.NewDecoder(f).Decode(&rules)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode rules file: %v", err)
	}

	return rules, nil
}
