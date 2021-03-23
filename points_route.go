package main

import (
	"fmt"
	"math/big"
	"net/http"
	"sort"
)

func (s server) handlePointsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		x := big.NewInt(0)
		y := big.NewInt(0)
		distance := big.NewInt(0)

		if _, ok := x.SetString(params.Get("x"), 10); ok == false || x.IsInt64() == false {
			err := messageResponse{
				Status:  400,
				Message: "invalid value for x axis",
			}
			s.respond(w, r, err, 400)
			return
		}

		if _, ok := y.SetString(params.Get("y"), 10); ok == false || y.IsInt64() == false {
			err := messageResponse{
				Status:  400,
				Message: "invalid value for x axis",
			}
			s.respond(w, r, err, 400)
			return
		}

		if _, ok := distance.SetString(params.Get("distance"), 10); ok == false || distance.IsInt64() == false {
			err := messageResponse{
				Status:  400,
				Message: "invalid value for distance",
			}
			s.respond(w, r, err, 400)
			return
		}

		if distance.Int64() < 0 {
			err := messageResponse{
				Status:  400,
				Message: "distance param can't be negative",
			}
			s.respond(w, r, err, 400)
			return
		}

		pointOrigin := point{
			X: x.Int64(),
			Y: y.Int64(),
		}

		// handling distances as strings to support numbers greater than MaxInt64
		pointsNearBy := s.pointsNearBy(pointOrigin, distance.String())

		response := pointsResponse{
			Count:  len(pointsNearBy),
			Points: pointsNearBy,
			Status: 200,
		}

		s.respond(w, r, response, 200)
	}
}

func (s server) pointsNearBy(origin point, distance string) []pointDistance {
	results := make([]pointDistance, 0)

	distance = fmt.Sprintf("%020s", distance)

	for _, current := range *s.data {
		// handling distances as leading 0 strings to make lessOrEqual operator work correctly
		if d := current.distanceFrom(origin); fmt.Sprintf("%020s", d.String()) <= distance {
			pd := pointDistance{
				Point:    current,
				Distance: d.String(),
			}
			results = append(results, pd)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		// handling distances as leading 0 strings to make less operator work correctly
		return fmt.Sprintf("%020s", results[i].Distance) < fmt.Sprintf("%020s", results[j].Distance)
	})

	return results
}
