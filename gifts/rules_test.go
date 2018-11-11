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
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				path: "testdata/invalidRules.json",
			},
			want:    &Rules{},
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
