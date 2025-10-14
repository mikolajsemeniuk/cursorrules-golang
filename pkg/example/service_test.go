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
	find   func(string) (string, error)
	update func(string, string) error
}

func (m *mock) Find(name string) (string, error)       { return m.find(name) }
func (m *mock) Update(name string, value string) error { return m.update(name, value) }

func TestServer_Find(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		wantBody := `{ "message": "ok" }`
		wantStatus := http.StatusOK
		m := &mock{find: func(string) (string, error) { return wantBody, nil }}
		recorder := httptest.NewRecorder()
		server := example.NewServer(m)

		request, err := http.NewRequest("GET", "/find?name=mike", nil)
		assert.NoError(t, err)
		server.Find(recorder, request)

		assert.Equal(t, wantStatus, recorder.Code)
		assert.Equal(t, wantBody, recorder.Body.String())
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		wantBody := `{ "message": "not found" }`
		wantStatus := http.StatusNotFound
		m := &mock{find: func(string) (string, error) { return wantBody, errors.New("not found") }}
		server := example.NewServer(m)
		recorder := httptest.NewRecorder()

		request, err := http.NewRequest("GET", "/find?name=mike", nil)
		assert.NoError(t, err)
		server.Find(recorder, request)

		assert.Equal(t, wantStatus, recorder.Code)
		assert.Equal(t, wantBody, recorder.Body.String())
	})
}

func TestServer_Update(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		wantBody := `{ "message": "success" }`
		wantStatus := http.StatusOK
		m := &mock{update: func(string, string) error { return nil }}
		server := example.NewServer(m)
		recorder := httptest.NewRecorder()

		request, _ := http.NewRequest("GET", "/update?name=mike&value=123", nil)
		server.Update(recorder, request)

		assert.Equal(t, wantStatus, recorder.Code)
		assert.Equal(t, wantBody, recorder.Body.String())
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		wantBody := `{ "message": "not found" }`
		wantStatus := http.StatusNotFound
		m := &mock{update: func(string, string) error { return errors.New("not found") }}
		server := example.NewServer(m)
		recorder := httptest.NewRecorder()

		request, _ := http.NewRequest("GET", "/update?name=mike&value=123", nil)
		server.Update(recorder, request)

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
