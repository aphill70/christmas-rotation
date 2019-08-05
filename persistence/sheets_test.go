package persistence

import (
	"reflect"
	"testing"

	sheets "google.golang.org/api/sheets/v4"
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

func TestSheet_getColumnToWrite(t *testing.T) {
	type fields struct {
		client    *sheets.Service
		cellRange string
		sheetID   string
		verbose   bool
		columnMap map[int]string
		rowMap    map[int]string
		lastRow   int
	}
	type args struct {
		year string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "new/column",
			fields: fields{
				client:    nil,
				cellRange: "",
				sheetID:   "",
				verbose:   false,

				columnMap: map[int]string{
					0: "Name",
					1: "2003",
					2: "2004",
				},
				rowMap:  map[int]string{},
				lastRow: 10,
			},
			args: args{
				year: "2005",
			},
			want: "D1:D11",
		},
		{
			name: "existing/column",
			fields: fields{
				client:    nil,
				cellRange: "",
				sheetID:   "",
				verbose:   false,

				columnMap: map[int]string{
					0: "Name",
					1: "2003",
					2: "2004",
				},
				rowMap:  map[int]string{},
				lastRow: 10,
			},
			args: args{
				year: "2003",
			},
			want: "B1:B11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sheet{
				client:    tt.fields.client,
				cellRange: tt.fields.cellRange,
				sheetID:   tt.fields.sheetID,
				verbose:   tt.fields.verbose,
				columnMap: tt.fields.columnMap,
				rowMap:    tt.fields.rowMap,
				lastRow:   tt.fields.lastRow,
			}
			if got := s.getColumnToWrite(tt.args.year); got != tt.want {
				t.Errorf("Sheet.getColumnToWrite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSheet_formatDataToWrite(t *testing.T) {
	type fields struct {
		client    *sheets.Service
		cellRange string
		sheetID   string
		verbose   bool
		columnMap map[int]string
		rowMap    map[int]string
		lastRow   int
	}
	type args struct {
		assignments map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]interface{}
	}{
		{
			name: "valid",
			fields: fields{
				lastRow: 5,
				rowMap: map[int]string{
					0: "jane",
					1: "pete",
					3: "jill",
					5: "bob",
				},
			},
			args: args{
				assignments: map[string]string{
					"jane": "pete",
					"pete": "jill",
					"jill": "bob",
					"bob":  "jane",
				},
			},
			want: [][]interface{}{
				[]interface{}{
					"pete",
					"jill",
					"",
					"bob",
					"",
					"jane",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sheet{
				client:    tt.fields.client,
				cellRange: tt.fields.cellRange,
				sheetID:   tt.fields.sheetID,
				verbose:   tt.fields.verbose,
				columnMap: tt.fields.columnMap,
				rowMap:    tt.fields.rowMap,
				lastRow:   tt.fields.lastRow,
			}
			if got := s.formatDataToWrite(tt.args.assignments); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sheet.formatDataToWrite() = %v, want %v", got, tt.want)
			}
		})
	}
}
