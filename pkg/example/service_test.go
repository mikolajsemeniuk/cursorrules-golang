package example_test

import (
	"cursor-rules-golang/pkg/example"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mock struct {
	find func() (string, error)
	add  func(string) error
}

func (m *mock) Find() (string, error) { return m.find() }
func (m *mock) Add(v string) error    { return m.add(v) }

func TestServer(t *testing.T) {
	t.Parallel()

	t.Run("FindOk", func(t *testing.T) {
		t.Parallel()

		wantBody := `{ "message": "ok" }`
		wantStatus := http.StatusOK
		m := &mock{find: func() (string, error) { return wantBody, nil }}
		recorder := httptest.NewRecorder()
		server := example.NewServer(m)

		server.Find(recorder, &http.Request{})

		assert.Equal(t, wantStatus, recorder.Code)
		assert.Equal(t, wantBody, recorder.Body.String())
	})

	t.Run("FindNotFound", func(t *testing.T) {
		t.Parallel()

		wantBody := `{ "message": "not found" }`
		wantStatus := http.StatusNotFound
		m := &mock{find: func() (string, error) { return wantBody, errors.New("not found") }}
		server := example.NewServer(m)
		recorder := httptest.NewRecorder()

		server.Find(recorder, &http.Request{})

		assert.Equal(t, wantStatus, recorder.Code)
		assert.Equal(t, wantBody, recorder.Body.String())
	})

	t.Run("AddOk", func(t *testing.T) {
		t.Parallel()

		wantBody := `{ "message": "success" }`
		wantStatus := http.StatusOK
		m := &mock{add: func(string) error { return nil }}
		server := example.NewServer(m)
		recorder := httptest.NewRecorder()

		request, _ := http.NewRequest("GET", "/add?name=mike", nil)
		server.Add(recorder, request)

		assert.Equal(t, wantStatus, recorder.Code)
		assert.Equal(t, wantBody, recorder.Body.String())
	})
}

func TestFormat(t *testing.T) {
	t.Parallel()

	t.Run("FormatValue", func(t *testing.T) {
		t.Parallel()

		in := "value"
		want := fmt.Sprintf("formatted: %s", in)
		got, err := example.Format(in)

		assert.NoError(t, err)
		assert.Equal(t, want, got, "The two words should be the same")
	})

	t.Run("ReturnError", func(t *testing.T) {
		t.Parallel()

		var in string
		var want string
		got, err := example.Format(in)

		assert.EqualError(t, err, "value cannot be empty string")
		assert.Equal(t, want, got, "The two words should be the same")
	})
}
