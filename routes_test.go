package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_routes_notFoundHandler(t *testing.T) {
	t.Run("return right statusCode", func(t *testing.T) {
		server := newServer()
		server.configureRoutes()

		req, err := http.NewRequest(http.MethodGet, "/api/invalidEndpoint", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		server.router.ServeHTTP(rr, req)

		if status, want := rr.Code, http.StatusNotFound; status != want {
			t.Errorf("handler returned wrong status code: got %d want %d", status, want)
		}
	})

	t.Run("return proper headers", func(t *testing.T) {
		server := newServer()
		server.configureRoutes()

		req, err := http.NewRequest(http.MethodGet, "/api/invalidEndpoint", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		server.router.ServeHTTP(rr, req)

		if got, want := rr.Header().Get("Content-Type"), "application/json"; !strings.Contains(got, want) {
			t.Errorf("handler returned wrong content-type: got '%v' want '%v'", got, want)
		}
	})

	t.Run("return valid json payload", func(t *testing.T) {
		server := newServer()
		server.configureRoutes()

		req, err := http.NewRequest(http.MethodGet, "/api/invalidEndpoint", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		server.router.ServeHTTP(rr, req)
		decoder := json.NewDecoder(rr.Body)

		var result messageResponse

		if err := decoder.Decode(&result); err != nil {
			t.Fatalf("error while decoding body %s", err)
		}

		if want := http.StatusNotFound; result.Status != want {
			t.Errorf("result message got: '%v' want: '%v'", result.Status, want)
		}

		if want := "Not Found"; result.Message != want {
			t.Errorf("result message got: '%v' want: '%v'", result.Message, want)
		}
	})
}

func Test_routes_methodNotAllowed(t *testing.T) {
	t.Run("return right statusCode", func(t *testing.T) {
		server := newServer()
		server.configureRoutes()

		req, err := http.NewRequest(http.MethodPost, "/api/points", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		server.router.ServeHTTP(rr, req)

		if got, want := rr.Code, http.StatusMethodNotAllowed; got != want {
			t.Errorf("handler returned wrong got code: got %d want %d", got, want)
		}
	})

	t.Run("return proper headers", func(t *testing.T) {
		server := newServer()
		server.configureRoutes()

		req, err := http.NewRequest(http.MethodPost, "/api/points", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		server.router.ServeHTTP(rr, req)

		if got, want := rr.Header().Get("Content-Type"), "application/json"; !strings.Contains(got, want) {
			t.Errorf("handler returned wrong content-type: got '%v' want '%v'", got, want)
		}
	})

	t.Run("return valid json payload", func(t *testing.T) {
		server := newServer()
		server.configureRoutes()

		req, err := http.NewRequest(http.MethodPost, "/api/points", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		server.router.ServeHTTP(rr, req)
		decoder := json.NewDecoder(rr.Body)

		var result messageResponse

		if err := decoder.Decode(&result); err != nil {
			t.Fatalf("error while decoding body %s", err)
		}

		if got, want := result.Status, http.StatusMethodNotAllowed; got != want {
			t.Errorf("result message got: '%v' want: '%v'", got, want)
		}

		if got, want := result.Message, "Method Not Allowed"; got != want {
			t.Errorf("result message got: '%v' want: '%v'", got, want)
		}
	})
}
