package gifts

import (
	"fmt"
	"testing"
)

var addGiverTests = []struct {
	recipient string
	giver     string
	year      string
}{
	{"Person1", "Giver1", "2009"},
}

var newGiftTests = []struct {
	in        string
	recipient string

	err bool
}{
	{"Person1", "Person1", false},
	{"", "", true},
}

func TestAddGiver(t *testing.T) {
	for _, tt := range addGiverTests {
		t.Run(fmt.Sprintf("%s/%s", tt.recipient, tt.giver), func(t *testing.T) {
			sut, _ := NewGift(tt.recipient)
			if sut != nil && sut.Recipient != tt.recipient {
				t.Errorf("got %q, want %q", sut.Recipient, tt.recipient)
			}

			if sut != nil && len(sut.Givers) != 0 {
				t.Error("Didn't initialize Gift correctly")
			}

			sut.AddGiver(tt.giver, tt.year)

			if sut != nil && tt.giver != "" && len(sut.Givers) == 0 {
				t.Error("Giver was not added correctly")
			}
		})
	}
}

func TestNewGift(t *testing.T) {
	for _, tt := range newGiftTests {
		t.Run(tt.in, func(t *testing.T) {
			sut, err := NewGift(tt.in)
			if (err == nil) == tt.err {
				t.Fail()
			}

			if sut != nil && sut.Recipient != tt.recipient {
				t.Errorf("got %q, want %q", sut.Recipient, tt.recipient)
			}
		})
	}
}
