package gifts

import (
	"reflect"
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

func TestRotation_GetEligibleGivers(t *testing.T) {
	type fields struct {
		RecipientToGiver map[string]string
		Recipients       map[string]*Gift
		Members          map[string]bool
		Rules            *Rules
		currentRecipient *Gift
	}
	type args struct {
		recipient string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]bool
		wantErr bool
	}{
		{
			name: "simple",
			fields: fields{
				RecipientToGiver: map[string]string{},
				Recipients: map[string]*Gift{
					"jill": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "bill",
							"2003": "jane",
						},
						Givers: map[string]bool{
							"bill": true,
							"jane": true,
						},
					},
					"bill": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "jeff",
							"2003": "jen",
						},
						Givers: map[string]bool{
							"jen":  true,
							"jeff": true,
						},
					},
					"jen": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "jill",
							"2003": "jeff",
						},
						Givers: map[string]bool{
							"jill": true,
							"jeff": true,
						},
					},
					"jeff": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "jen",
							"2003": "jill",
						},
						Givers: map[string]bool{
							"jen":  true,
							"jill": true,
						},
					},
				},
				Members: map[string]bool{
					"jeff": true,
					"jen":  true,
					"jill": true,
					"bill": true,
				},
				Rules:            &Rules{},
				currentRecipient: &Gift{},
			},
			args: args{
				recipient: "jeff",
			},
			want: map[string]bool{
				"bill": true,
			},
			wantErr: false,
		},
		{
			name: "rollover/initial/full",
			fields: fields{
				RecipientToGiver: map[string]string{},
				Recipients: map[string]*Gift{
					"bill": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "jane",
							"2003": "jen",
						},
						Givers: map[string]bool{
							"jane": true,
							"jen":  true,
						},
					},
					"jen": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "bill",
							"2003": "jane",
						},
						Givers: map[string]bool{
							"bill": true,
							"jane": true,
						},
					},
					"jane": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "jane",
							"2003": "bill",
						},
						Givers: map[string]bool{
							"jane": true,
							"bill": true,
						},
					},
				},
				Members: map[string]bool{
					"jane": true,
					"jen":  true,
					"bill": true,
				},
				Rules:            &Rules{},
				currentRecipient: &Gift{},
			},
			args: args{
				recipient: "jen",
			},
			want: map[string]bool{
				"bill": true,
				"jane": true,
			},
			wantErr: false,
		},
		{
			name: "rollover/repeat/full",
			fields: fields{
				RecipientToGiver: map[string]string{},
				Recipients: map[string]*Gift{
					"bill": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "jane",
							"2003": "jen",
							"2004": "jen",
						},
						Givers: map[string]bool{
							"jen": true,
						},
					},
					"jen": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "bill",
							"2003": "jane",
						},
						Givers: map[string]bool{
							"bill": true,
							"jane": true,
						},
					},
					"jane": &Gift{
						Recipient: "",
						History: map[string]string{
							"2002": "jane",
							"2003": "bill",
						},
						Givers: map[string]bool{
							"jane": true,
							"bill": true,
						},
					},
				},
				Members: map[string]bool{
					"jane": true,
					"jen":  true,
					"bill": true,
				},
				Rules:            &Rules{},
				currentRecipient: &Gift{},
			},
			args: args{
				recipient: "bill",
			},
			want: map[string]bool{
				"jane": true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotation{
				RecipientToGiver: tt.fields.RecipientToGiver,
				Recipients:       tt.fields.Recipients,
				Members:          tt.fields.Members,
				Rules:            tt.fields.Rules,
				currentRecipient: tt.fields.currentRecipient,
			}
			got, err := r.GetEligibleGivers(tt.args.recipient)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rotation.GetEligibleGivers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotation.GetEligibleGivers() = %v, want %v", got, tt.want)
			}
		})
	}
}
