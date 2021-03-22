package main

import (
	"encoding/json"
	"net/http"
)

func (s *server) configureRoutes() {
	s.router.NotFoundHandler = notFoundHandler()
	s.router.MethodNotAllowedHandler = methodNotAllowedHandler()

	s.router.HandleFunc("/points", s.handlePointsGet()).Methods("GET")
}

func notFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(messageResponse{
			Status:  http.StatusNotFound,
			Message: "Not Found",
		})
	}
}

func methodNotAllowedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(messageResponse{
			Status:  http.StatusMethodNotAllowed,
			Message: "Method Not Allowed",
		})
	}
}
