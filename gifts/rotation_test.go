package gifts

import (
	"fmt"
	"testing"
)

var addRecipientTests = []struct {
	recipient string
}{
	{"Person1"},
	{"Person2"},
	{"Person3"},
}

func TestAddRecipeint(t *testing.T) {
	for _, tt := range addRecipientTests {
		t.Run(fmt.Sprintf("%s", tt.recipient), func(t *testing.T) {
			sut, _ := NewRotation()

			sut.AddRecipient(tt.recipient)

			if sut != nil && sut.currentRecipient != nil && sut.currentRecipient.Recipient != tt.recipient {
				t.Errorf("got %q, want %q", sut.currentRecipient.Recipient, tt.recipient)
			}

			if sut != nil && !sut.allMembers[tt.recipient] {
				t.Errorf("recipient not found in the all members: %s", tt.recipient)
			}
		})
	}
}

func TestRotation_AddGiver(t *testing.T) {
	type fields struct {
		recipientGiver   map[string]string
		Recipients       map[string]*Gift
		currentRecipient *Gift
		allMembers       map[string]bool
	}
	type args struct {
		giver string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotation{
				recipientGiver:   tt.fields.recipientGiver,
				Recipients:       tt.fields.Recipients,
				currentRecipient: tt.fields.currentRecipient,
				allMembers:       tt.fields.allMembers,
			}
			if err := r.AddGiver(tt.args.giver); (err != nil) != tt.wantErr {
				t.Errorf("Rotation.AddGiver() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
