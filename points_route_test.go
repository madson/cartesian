package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_server_handlePointsGet(t *testing.T) {
	t.Run("handlePointsGet no URL params", func(t *testing.T) {
		server := newServer()

		req, err := http.NewRequest(http.MethodGet, "/api/points", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		server.handlePointsGet().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("handlePointsGet invalid value for x axis", func(t *testing.T) {
		server := newServer()

		req, err := http.NewRequest(http.MethodGet, "/api/points", nil)
		if err != nil {
			t.Fatal(err)
		}

		tests := []string{
			"invalid",
			"*",
			" ",
		}

		for _, value := range tests {
			q := req.URL.Query()
			q.Add("x", value)
			q.Add("y", "0")
			q.Add("distance", "1000")
			req.URL.RawQuery = q.Encode()

			rr := httptest.NewRecorder()
			server.handlePointsGet().ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
			}
		}
	})

	t.Run("handlePointsGet invalid value for y axis", func(t *testing.T) {
		server := newServer()

		req, err := http.NewRequest(http.MethodGet, "/api/points", nil)
		if err != nil {
			t.Fatal(err)
		}

		tests := []string{
			"invalid",
			"*",
			" ",
		}

		for _, value := range tests {
			q := req.URL.Query()
			q.Add("x", "0")
			q.Add("y", value)
			q.Add("distance", "1000")
			req.URL.RawQuery = q.Encode()

			rr := httptest.NewRecorder()
			server.handlePointsGet().ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
			}
		}
	})

	t.Run("handlePointsGet invalid value for distance", func(t *testing.T) {
		server := newServer()

		req, err := http.NewRequest(http.MethodGet, "/api/points", nil)
		if err != nil {
			t.Fatal(err)
		}

		tests := []string{
			"invalid",
			"*",
			" ",
		}

		for _, value := range tests {
			q := req.URL.Query()
			q.Add("x", "0")
			q.Add("y", "0")
			q.Add("distance", value)
			req.URL.RawQuery = q.Encode()

			rr := httptest.NewRecorder()
			server.handlePointsGet().ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
			}
		}
	})
}

func Test_server_pointsNearBy(t *testing.T) {

	t.Run("pointsNearBy results count", func(t *testing.T) {
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
