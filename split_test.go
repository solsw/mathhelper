package mathhelper

import (
	"reflect"
	"testing"
)

func TestSplit_int(t *testing.T) {
	type args struct {
		n     int
		parts int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "01",
			args:    args{n: 0, parts: 1},
			wantErr: true,
		},
		{name: "02",
			args:    args{n: 8, parts: 0},
			wantErr: true,
		},
		{name: "03",
			args:    args{n: 8, parts: 9},
			wantErr: true,
		},
		{name: "1",
			args: args{n: 8, parts: 1},
			want: []int{8},
		},
		{name: "2",
			args: args{n: 8, parts: 2},
			want: []int{4, 4},
		},
		{name: "3",
			args: args{n: 8, parts: 3},
			want: []int{3, 3, 2},
		},
		{name: "4",
			args: args{n: 8, parts: 4},
			want: []int{2, 2, 2, 2},
		},
		{name: "5",
			args: args{n: 8, parts: 5},
			want: []int{2, 2, 2, 1, 1},
		},
		{name: "6",
			args: args{n: 8, parts: 6},
			want: []int{2, 2, 1, 1, 1, 1},
		},
		{name: "7",
			args: args{n: 8, parts: 7},
			want: []int{2, 1, 1, 1, 1, 1, 1},
		},
		{name: "8",
			args: args{n: 8, parts: 8},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Split(tt.args.n, tt.args.parts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
