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
				Recipient: "bob",
				History:   map[string]string{},
				Givers:    map[string]bool{},
			},
			wantErr: false,
		},
		{
			name: "chris/valid",
			args: args{
				recipient: "chris",
			},
			want: &Gift{
				Recipient: "christopher",
				History:   map[string]string{},
				Givers:    map[string]bool{},
			},
			wantErr: false,
		},
		{
			name: "christopher/valid",
			args: args{
				recipient: "christopher",
			},
			want: &Gift{
				Recipient: "christopher",
				History:   map[string]string{},
				Givers:    map[string]bool{},
			},
			wantErr: false,
		},
		{
			name: "micheal/valid",
			args: args{
				recipient: "micheal",
			},
			want: &Gift{
				Recipient: "michael",
				History:   map[string]string{},
				Givers:    map[string]bool{},
			},
			wantErr: false,
		},
		{
			name: "michael/valid",
			args: args{
				recipient: "michael",
			},
			want: &Gift{
				Recipient: "michael",
				History:   map[string]string{},
				Givers:    map[string]bool{},
			},
			wantErr: false,
		},
		{
			name: "michael/space/valid",
			args: args{
				recipient: "michael ",
			},
			want: &Gift{
				Recipient: "michael",
				History:   map[string]string{},
				Givers:    map[string]bool{},
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

func TestGift_AddGiver(t *testing.T) {
	type fields struct {
		Recipient string
		History   map[string]string
		Givers    map[string]bool
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
			name: "valid",
			fields: fields{
				Recipient: "bob",
				History:   map[string]string{"2002": "jane", "2003": "bill"},
				Givers:    map[string]bool{"jane": true, "bill": true},
			},
			args:    args{giver: "jill", year: "2004"},
			wantErr: false,
		},
		{
			name: "repeatYear",
			fields: fields{
				Recipient: "bob",
				History:   map[string]string{"2002": "jane", "2003": "bill"},
				Givers:    map[string]bool{"jane": true, "bill": true},
			},
			args:    args{giver: "jill", year: "2002"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gift{
				Recipient: tt.fields.Recipient,
				History:   tt.fields.History,
				Givers:    tt.fields.Givers,
			}
			err := g.AddGiver(tt.args.giver, tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGift() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
