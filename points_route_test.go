package main

import (
	"fmt"
	"testing"
)

func Test_server_handlePointsGet(t *testing.T) {

}

func Test_server_pointsNearBy(t *testing.T) {

	t.Run("pointsNearBy len", func(t *testing.T) {
		server := newServer()
		*server.data = []point{
			{X: -59, Y: 38},
			{X: 2, Y: -8},
			{X: -5, Y: 95},
			{X: -6, Y: -10},
			{X: 18, Y: -3},
			{X: -20, Y: 2},
			{X: -9, Y: -12},
		}

		origin := point{X: 1, Y: 1}

		tests := []struct {
			distance int
			want     int
		}{
			{10, 1},
			{18, 2},
			{21, 3},
			{22, 4},
			{23, 5},
			{24, 5},
			{100, 7},
		}

		for _, tt := range tests {
			pointsNearBy := server.pointsNearBy(origin, fmt.Sprint(tt.distance))

			if count := len(pointsNearBy); count != tt.want {
				t.Errorf("pointsNearBy() want count of %d, but got %d results", tt.want, count)
			}

		}
	})
}
