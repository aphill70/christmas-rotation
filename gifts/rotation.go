package gifts

import (
	"fmt"
	"sort"
)

type (
	// Rotation represents the gift rotation
	Rotation struct {
		RecipientToGiver map[string]string
		Recipients       map[string]*Gift
		Members          map[string]bool

		LastRow int

		Rules *Rules

		currentRecipient *Gift
	}
)

// Recipient -> [Giver : Count]

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
	recipient = NormalizeName(recipient)
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
	giver = NormalizeName(giver)
	if r.currentRecipient == nil {
		return fmt.Errorf("current recipient is null")
	}

	r.currentRecipient.AddGiver(giver, year)

	// could this be an int to handle rollover especially uneven rollover when we need to rollover and still remember?
	if !r.Members[giver] {
		r.Members[giver] = true
	}

	return nil
}

// GetEligibleGivers returns all valid givers for a given recipient
func (r *Rotation) GetEligibleGivers(recipient string) (map[string]int, error) {
	fmt.Println("TEST: ", recipient)
	recipient = NormalizeName(recipient)
	if !r.Members[recipient] || r.Recipients[recipient] == nil {
		return nil, fmt.Errorf("Invalid Recipient: %s", recipient)
	}

	gift := r.Recipients[recipient]
	var eligibleMembers = make(map[string]int)
	for member := range r.Members {
		// fmt.Printf("%v == %v\n", member, recipient)
		if member == recipient || !r.Rules.IsAllowed(recipient, member) {
			continue
		}

		giver := gift.GetGiver(member)
		if giver == nil {
			fmt.Println("New Giver: ", member)
			eligibleMembers[member] = 0
		} else {
			eligibleMembers[member] = giver.Count
		}
	}
	return eligibleMembers, nil
}

type rotationOptions struct {
	recipient   string
	lowestIndex int
	indexes     []int

	options map[int][]string
}

// ByEligibleGiversCount implements sort interface for sorting by the number of potential givers
type ByEligibleGiversCount []rotationOptions

func (e ByEligibleGiversCount) Len() int      { return len(e) }
func (e ByEligibleGiversCount) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e ByEligibleGiversCount) Less(i, j int) bool {
	lowestI := getLowestOptionIndex(e, i)
	lowestJ := getLowestOptionIndex(e, j)

	if lowestI != lowestJ {
		return lowestI < lowestJ
	}

	return len(e[i].options[lowestI]) < len(e[j].options[lowestJ])
}

func getLowestOptionIndex(e []rotationOptions, i int) int {
	lowestI := -1
	for iKey := range e[i].options {
		if lowestI == -1 || iKey < lowestI {
			lowestI = iKey
		}
	}

	return lowestI
}

// GetNextYearsRotation chooses next years givers based on rules and previous years
func (r *Rotation) GetNextYearsRotation(year string) {
	// Givers sorted by rotation index
	// Then members sorted by givers in current rotation index
	var options = []rotationOptions{}

	for member := range r.Members {
		optionList, _ := r.GetEligibleGivers(member)
		sortedOptions := map[int][]string{}
		lowestIndex := -1
		indexes := []int{}
		for option, count := range optionList {
			if sortedOptions[count] == nil {
				if lowestIndex == -1 || count < lowestIndex {
					lowestIndex = count
				}
				indexes = append(indexes, count)
				sortedOptions[count] = []string{}
			}

			sortedOptions[count] = append(sortedOptions[count], option)
		}

		sort.Ints(indexes)
		options = append(options, rotationOptions{
			recipient:   member,
			lowestIndex: lowestIndex,
			indexes:     indexes,
			options:     sortedOptions,
		})
	}

	sort.Sort(ByEligibleGiversCount(options))

	fmt.Println("RESULTS: ", options)

	used := map[string]bool{}
	for _, part := range options {
		found := false
		for currentIndex := range part.indexes {
			for _, option := range part.options[currentIndex] {
				if !used[option] {
					used[option] = true
					r.RecipientToGiver[part.recipient] = option
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
}
