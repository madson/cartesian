package main

import (
	_ "embed"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed data/points.json
var pointsData []byte

type server struct {
	router *mux.Router
	data   *[]point
}

type messageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type pointsResponse struct {
	Count  int             `json:"count"`
	Status int             `json:"status"`
	Points []pointDistance `json:"points"`
}

func newServer() *server {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	return &server{
		router: router,
		data:   &[]point{},
	}
}

func (s *server) respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	if data != nil {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(messageResponse{
		Status:  http.StatusNoContent,
		Message: "No Content",
	})
}

func (s *server) loadData() error {
	var records []point
	if err := json.Unmarshal(pointsData, &records); err != nil {
		return err
	}
	*s.data = records
	return nil
}
