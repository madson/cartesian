package main

import (
	"net/http"
	"sort"
	"strconv"
)

func (s server) handlePointsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		x, err := strconv.ParseInt(params.Get("x"), 10, 64)
		if err != nil {
			error := messageResponse{
				Status:  400,
				Message: "invalid URL query param: x",
			}
			s.respond(w, r, error, 400)
			return
		}

		y, err := strconv.ParseInt(params.Get("y"), 10, 64)
		if err != nil {
			error := messageResponse{
				Status:  400,
				Message: "invalid URL query param: y",
			}
			s.respond(w, r, error, 400)
			return
		}

		distance, err := strconv.ParseInt(params.Get("distance"), 10, 64)
		if err != nil {
			error := messageResponse{
				Status:  400,
				Message: "invalid URL query param: distance",
			}
			s.respond(w, r, error, 400)
			return
		}

		if distance < 0 {
			error := messageResponse{
				Status:  400,
				Message: "distance param can't be negative",
			}
			s.respond(w, r, error, 400)
			return
		}

		pointOrigin := point{
			X: x,
			Y: y,
		}

		pointsNearBy := s.pointsNearBy(pointOrigin, distance)

		response := pointsResponse{
			Count:  len(pointsNearBy),
			Points: pointsNearBy,
			Status: 200,
		}

		s.respond(w, r, response, 200)
	}
}

func (s server) pointsNearBy(origin point, distance int64) []pointDistance {
	results := make([]pointDistance, 0)

	for _, current := range *s.data {
		if d := current.distanceFrom(origin); d <= distance {
			pd := pointDistance{
				Point:    current,
				Distance: d,
			}
			results = append(results, pd)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})

	return results
}
