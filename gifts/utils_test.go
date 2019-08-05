package gifts

import "testing"

func Test_normalizeName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bob",
			args: args{
				name: "Bob",
			},
			want: "bob",
		},
		{
			name: "HILL",
			args: args{
				name: "HILL",
			},
			want: "hill",
		},
		{
			name: "Christopher",
			args: args{
				name: "Christopher",
			},
			want: "christopher",
		},
		{
			name: "Chris",
			args: args{
				name: "Chris",
			},
			want: "christopher",
		},
		{
			name: "Michael",
			args: args{
				name: "Michael",
			},
			want: "michael",
		},
		{
			name: "Micheal",
			args: args{
				name: "Micheal",
			},
			want: "michael",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeName(tt.args.name); got != tt.want {
				t.Errorf("normalizeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
