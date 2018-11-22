package persistence

import (
	"testing"
)

func Test_letterToColumn(t *testing.T) {
	type args struct {
		letter string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A",
			args: args{
				letter: "A",
			},
			want: 1,
		},
		{
			name: "AA",
			args: args{
				letter: "AA",
			},
			want: 27,
		},
		{
			name: "B",
			args: args{
				letter: "B",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterToColumn(tt.args.letter); got != tt.want {
				t.Errorf("letterToColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_columnToLetter(t *testing.T) {
	type args struct {
		column int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "A",
			args: args{
				column: 1,
			},
			want: "A",
		},
		{
			name: "AA",
			args: args{
				column: 27,
			},
			want: "AA",
		},
		{
			name: "B",
			args: args{
				column: 2,
			},
			want: "B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := columnToLetter(tt.args.column); got != tt.want {
				t.Errorf("columnToLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
