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
					[]string{"donna", "michael", "austin"},
					[]string{"megan", "christopher", "aiden", "ashlyn", "kenzie"},
					[]string{"heidi", "spencer", "charlie", "sadie"},
					[]string{"pop", "gigi"},
					[]string{"adam", "kaitie", "seeley"},
				},
				PersonLookup: map[string]map[string]bool{
					"donna": map[string]bool{
						"michael": true,
						"austin":  true,
					},
					"michael": map[string]bool{
						"austin": true,
						"donna":  true,
					},
					"austin": map[string]bool{
						"michael": true,
						"donna":   true,
					},
					"megan": map[string]bool{
						"christopher": true,
						"aiden":       true,
						"ashlyn":      true,
						"kenzie":      true,
					},
					"christopher": map[string]bool{
						"megan":  true,
						"aiden":  true,
						"ashlyn": true,
						"kenzie": true,
					},
					"aiden": map[string]bool{
						"megan":       true,
						"christopher": true,
						"ashlyn":      true,
						"kenzie":      true,
					},
					"ashlyn": map[string]bool{
						"megan":       true,
						"christopher": true,
						"aiden":       true,
						"kenzie":      true,
					},
					"kenzie": map[string]bool{
						"megan":       true,
						"christopher": true,
						"aiden":       true,
						"ashlyn":      true,
					},
					"heidi": map[string]bool{
						"spencer": true,
						"charlie": true,
						"sadie":   true,
					},
					"spencer": map[string]bool{
						"heidi":   true,
						"charlie": true,
						"sadie":   true,
					},
					"charlie": map[string]bool{
						"heidi":   true,
						"spencer": true,
						"sadie":   true,
					},
					"sadie": map[string]bool{
						"heidi":   true,
						"spencer": true,
						"charlie": true,
					},
					"pop": map[string]bool{
						"gigi": true,
					},
					"gigi": map[string]bool{
						"pop": true,
					},
					"adam": map[string]bool{
						"kaitie": true,
						"seeley": true,
					},
					"kaitie": map[string]bool{
						"adam":   true,
						"seeley": true,
					},
					"seeley": map[string]bool{
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
