package example

import (
	"errors"
	"fmt"
	"net/http"
)

type Storage interface {
	Find() (string, error)
	Add(string) error
}

type Server struct {
	storage Storage
}

func NewServer(s Storage) *Server {
	return &Server{storage: s}
}

func (s *Server) Find(w http.ResponseWriter, r *http.Request) {
	response, err := s.storage.Find()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{ "message": "not found" }`)
		return
	}

	fmt.Fprint(w, response)
}

func (s *Server) Add(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if err := s.storage.Add(name); err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{ "message": "not found" }`)
		return
	}

	fmt.Fprint(w, `{ "message": "success" }`)
}

func Format(v string) (string, error) {
	if v == "" {
		return "", errors.New("value cannot be empty string")
	}

	return fmt.Sprintf("formatted: %s", v), nil
}
