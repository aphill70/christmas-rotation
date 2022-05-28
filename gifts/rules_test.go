package gifts

import (
	"reflect"
	"testing"
)

func TestNewRules(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *Rules
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				path: "testdata/rules.json",
			},
			want: &Rules{
				Households: [][]string{
					{"donna", "michael", "austin"},
					{"megan", "christopher", "aiden", "ashlyn", "kenzie"},
					{"heidi", "spencer", "charlie", "sadie"},
					{"pop", "gigi"},
					{"adam", "kaitie", "seeley"},
				},
				PersonLookup: map[string]map[string]bool{
					"donna": {
						"michael": true,
						"austin":  true,
					},
					"michael": {
						"austin": true,
						"donna":  true,
					},
					"austin": {
						"michael": true,
						"donna":   true,
					},
					"megan": {
						"christopher": true,
						"aiden":       true,
						"ashlyn":      true,
						"kenzie":      true,
					},
					"christopher": {
						"megan":  true,
						"aiden":  true,
						"ashlyn": true,
						"kenzie": true,
					},
					"aiden": {
						"megan":       true,
						"christopher": true,
						"ashlyn":      true,
						"kenzie":      true,
					},
					"ashlyn": {
						"megan":       true,
						"christopher": true,
						"aiden":       true,
						"kenzie":      true,
					},
					"kenzie": {
						"megan":       true,
						"christopher": true,
						"aiden":       true,
						"ashlyn":      true,
					},
					"heidi": {
						"spencer": true,
						"charlie": true,
						"sadie":   true,
					},
					"spencer": {
						"heidi":   true,
						"charlie": true,
						"sadie":   true,
					},
					"charlie": {
						"heidi":   true,
						"spencer": true,
						"sadie":   true,
					},
					"sadie": {
						"heidi":   true,
						"spencer": true,
						"charlie": true,
					},
					"pop": {
						"gigi": true,
					},
					"gigi": {
						"pop": true,
					},
					"adam": {
						"kaitie": true,
						"seeley": true,
					},
					"kaitie": {
						"adam":   true,
						"seeley": true,
					},
					"seeley": {
						"adam":   true,
						"kaitie": true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				path: "testdata/invalidRules.json",
			},
			want: &Rules{
				PersonLookup: map[string]map[string]bool{},
			},
			wantErr: false,
		},
		{
			name: "invalid/json",
			args: args{
				path: "testdata/invalid.json",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid/missing",
			args: args{
				path: "testdata/missing.json",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRules(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRules() = %v, want %v", got, tt.want)
			}
		})
	}
}
