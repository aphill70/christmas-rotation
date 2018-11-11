package gifts

import (
	"testing"
)

func TestRotation_AddGiver(t *testing.T) {
	type fields struct {
		RecipientToGiver map[string]string
		Recipients       map[string]*Gift
		currentRecipient *Gift
		Members          map[string]bool
	}
	type args struct {
		giver string
		year  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid/currentRecipient",
			fields: fields{
				RecipientToGiver: map[string]string{},
				Recipients:       map[string]*Gift{},
				currentRecipient: &Gift{
					Recipient: "jane",
					History:   map[string]string{},
					Givers:    map[string]bool{},
				},
				Members: map[string]bool{},
			},
			args: args{
				giver: "bill",
				year:  "2018",
			},
			wantErr: false,
		},
		{
			name: "nil/currentRecipient",
			fields: fields{
				RecipientToGiver: map[string]string{},
				Recipients:       map[string]*Gift{},
				currentRecipient: nil,
				Members:          map[string]bool{},
			},
			args: args{
				giver: "test",
				year:  "2018",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotation{
				RecipientToGiver: tt.fields.RecipientToGiver,
				Recipients:       tt.fields.Recipients,
				currentRecipient: tt.fields.currentRecipient,
				Members:          tt.fields.Members,
			}
			if err := r.AddGiver(tt.args.giver, tt.args.year); (err != nil) != tt.wantErr {
				t.Errorf("Rotation.AddGiver() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
