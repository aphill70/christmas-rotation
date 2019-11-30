package gifts

import (
	"reflect"
	"testing"
)

func TestNewGift(t *testing.T) {
	type args struct {
		recipient string
	}
	tests := []struct {
		name    string
		args    args
		want    *Gift
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				recipient: "bob",
			},
			want: &Gift{
				Recipient:   "bob",
				History:     make(map[string]*Giver),
				Givers:      []*Giver{},
				GiverLookup: make(map[string]*Giver),
			},
			wantErr: false,
		},
		{
			name: "chris/valid",
			args: args{
				recipient: "chris",
			},
			want: &Gift{
				Recipient:   "christopher",
				History:     make(map[string]*Giver),
				Givers:      []*Giver{},
				GiverLookup: make(map[string]*Giver),
			},
			wantErr: false,
		},
		{
			name: "christopher/valid",
			args: args{
				recipient: "christopher",
			},
			want: &Gift{
				Recipient:   "christopher",
				History:     make(map[string]*Giver),
				Givers:      []*Giver{},
				GiverLookup: make(map[string]*Giver),
			},
			wantErr: false,
		},
		{
			name: "micheal/valid",
			args: args{
				recipient: "micheal",
			},
			want: &Gift{
				Recipient:   "michael",
				History:     make(map[string]*Giver),
				Givers:      []*Giver{},
				GiverLookup: make(map[string]*Giver),
			},
			wantErr: false,
		},
		{
			name: "michael/valid",
			args: args{
				recipient: "michael",
			},
			want: &Gift{
				Recipient:   "michael",
				History:     make(map[string]*Giver),
				Givers:      []*Giver{},
				GiverLookup: make(map[string]*Giver),
			},
			wantErr: false,
		},
		{
			name: "michael/space/valid",
			args: args{
				recipient: "michael ",
			},
			want: &Gift{
				Recipient:   "michael",
				History:     make(map[string]*Giver),
				Givers:      []*Giver{},
				GiverLookup: make(map[string]*Giver),
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				recipient: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGift(tt.args.recipient)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGift() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getGiftObject() Gift {
	jane := &Giver{
		Giver: "jane",
		Count: 2,
	}

	bill := &Giver{
		Giver: "bill",
		Count: 1,
	}

	jill := &Giver{
		Giver: "jill",
		Count: 1,
	}

	return Gift{
		Recipient: "bob",

		History: map[string]*Giver{
			"2002": jane,
			"2003": bill,
			"2004": jill,
			"2005": jane,
		},

		Givers: []*Giver{jill, bill, jane},

		GiverLookup: map[string]*Giver{
			"jane": jane,
			"bill": bill,
			"jill": jill,
		},
	}
}

func TestGift_AddGiver(t *testing.T) {
	type args struct {
		giver string
		year  string
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		wantGiverIndex string
		wantGiver      Giver
	}{
		{
			name:           "valid",
			args:           args{giver: "jim", year: "2006"},
			wantErr:        false,
			wantGiverIndex: "jim",
			wantGiver: Giver{
				Giver: "jim",
				Count: 1,
			},
		},
		{
			name:    "repeatYear",
			args:    args{giver: "jill", year: "2002"},
			wantErr: true,
		},
		{
			name:           "rollover/bill",
			args:           args{giver: "bill", year: "2006"},
			wantErr:        false,
			wantGiverIndex: "bill",
			wantGiver: Giver{
				Giver: "bill",
				Count: 2,
			},
		},
		{
			name:           "rollover/jane",
			args:           args{giver: "jane", year: "2006"},
			wantErr:        false,
			wantGiverIndex: "jane",
			wantGiver: Giver{
				Giver: "jane",
				Count: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := getGiftObject()

			err := g.AddGiver(tt.args.giver, tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGift() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			giverResult := g.GiverLookup[tt.wantGiverIndex]

			if !tt.wantErr && giverResult.Giver != tt.wantGiver.Giver {
				t.Errorf("%s %s", giverResult.Giver, tt.wantGiver.Giver)
				return
			}

			if !tt.wantErr && giverResult.Count != tt.wantGiver.Count {
				t.Errorf("%d%d", giverResult.Count, tt.wantGiver.Count)
				return
			}
		})
	}
}

func TestGift_GetCurrentRotationIndex(t *testing.T) {
	type fields struct {
		GiverLookup map[string]*Giver
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "3_different",
			fields: fields{
				GiverLookup: map[string]*Giver{
					"jill":  &Giver{Count: 1},
					"jen":   &Giver{Count: 2},
					"james": &Giver{Count: 3},
				},
			},
			want: 1,
		},
		{
			name: "2_matching",
			fields: fields{
				GiverLookup: map[string]*Giver{
					"jill":  &Giver{Count: 1},
					"jen":   &Giver{Count: 2},
					"james": &Giver{Count: 1},
				},
			},
			want: 1,
		},
		{
			name: "all_the_same",
			fields: fields{
				GiverLookup: map[string]*Giver{
					"jill":  &Giver{Count: 1},
					"jen":   &Giver{Count: 1},
					"james": &Giver{Count: 1},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gift{
				GiverLookup: tt.fields.GiverLookup,
			}
			if got := g.GetCurrentRotationIndex(); got != tt.want {
				t.Errorf("Gift.GetCurrentRotationIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
