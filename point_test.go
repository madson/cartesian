package main

import (
	"math"
	"math/big"
	"testing"
)

func Test_point_distanceFrom(t *testing.T) {
	min := int64(math.MinInt64)
	max := int64(math.MaxInt64)

	tests := []struct {
		origin point
		dest   point
		want   string
	}{
		{origin: point{X: 1, Y: 1}, dest: point{X: 1, Y: 3}, want: "2"},
		{origin: point{X: -1, Y: -1}, dest: point{X: -20, Y: -20}, want: "38"},
		{origin: point{X: min, Y: min}, dest: point{X: max, Y: max - 1}, want: "36893488147419103229"},
		{origin: point{X: 0, Y: 0}, dest: point{X: 222, Y: 222}, want: "444"},
		{origin: point{X: -10, Y: -10}, dest: point{X: 10, Y: 10}, want: "40"},
		{origin: point{X: -20, Y: -20}, dest: point{X: 10, Y: 10}, want: "60"},
		{origin: point{X: -999, Y: 0}, dest: point{X: 999, Y: 1}, want: "1999"},
		{origin: point{X: max, Y: 1}, dest: point{X: 0, Y: 2}, want: big.NewInt(0).Add(big.NewInt(max), big.NewInt(1)).String()},
		{origin: point{X: max, Y: max}, dest: point{X: min, Y: min}, want: "36893488147419103230"},
		{origin: point{X: max, Y: min}, dest: point{X: min, Y: max - 2}, want: "36893488147419103228"},
		{origin: point{X: min, Y: max}, dest: point{X: max, Y: min + 3}, want: "36893488147419103227"},
	}
	for _, tt := range tests {
		t.Run("distance_"+tt.want, func(t *testing.T) {
			p1 := tt.origin
			p2 := tt.dest
			if got := p1.distanceFrom(p2); got.String() != tt.want {
				t.Errorf("distanceFrom() got %v, want %v", got, tt.want)
			}
		})
	}
}
