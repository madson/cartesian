package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

var (
	port = flag.String("port", "8000", "server port number")
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	s := newServer()
	s.configureRoutes()

	err := s.loadData()
	if err != nil {
		return err
	}

	log.Println("server started on http://localhost:" + *port)

	err = http.ListenAndServe(":"+*port, handlers.CompressHandler(s.router))
	if err != nil {
		return err
	}

	return nil
}
