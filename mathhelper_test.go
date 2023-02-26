//go:build go1.18

package mathhelper

import (
	"math"
	"reflect"
	"testing"
)

func TestIsEven_int(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{i: 23}, want: false},
		{name: "2", args: args{i: 2332}, want: true},
		{name: "3", args: args{i: -23}, want: false},
		{name: "4", args: args{i: -2332}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.i); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEven_uint8(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{i: 23}, want: false},
		{name: "2", args: args{i: 32}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.i); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrac_float32(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: "1", args: args{f: 2.1}, want: 0.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Frac(tt.args.f); Abs(got-tt.want) > 0.0001 {
				t.Errorf("Frac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrac_float64(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "2", args: args{f: math.Pi}, want: 0.1415926},
		{name: "3", args: args{f: -math.E}, want: -0.718281828},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Frac(tt.args.f); Abs(got-tt.want) > 0.0001 {
				t.Errorf("Frac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntMin_int(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0"},
		{name: "1", args: args{m: 1, n: 2}, want: 1},
		{name: "2", args: args{m: -1, n: -2}, want: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntMin(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntMax_byte(t *testing.T) {
	type args struct {
		m byte
		n byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "0"},
		{name: "1", args: args{m: 1, n: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntMax(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
