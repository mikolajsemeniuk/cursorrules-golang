package example

import (
	"errors"
	"fmt"
	"net/http"
)

type Storage interface {
	Find(name string) (string, error)
	Update(name string, value string) error
}

type Server struct {
	storage Storage
}

func NewServer(s Storage) *Server {
	return &Server{storage: s}
}

func (s *Server) Find(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	res, err := s.storage.Find(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{ "message": "not found" }`)
		return
	}

	fmt.Fprint(w, res)
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	value := query.Get("value")

	if err := s.storage.Update(name, value); err != nil {
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
