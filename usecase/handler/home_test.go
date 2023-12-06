package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestNewHome(t *testing.T) {
	router := mux.NewRouter()

	resultRouter := NewHome(router)

	assert.Equal(t, router, resultRouter, "Router mismatch")
}

func TestHomeHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		fnParseFiles = func(filenames ...string) (*template.Template, error) {
			return template.New("mock").Parse("mock template")
		}
		homeHandler(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")

		expectedBody := "mock template"
		assert.Contains(t, rr.Body.String(), expectedBody, "Response body mismatch")
	})

	t.Run("failed parse files", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		fnParseFiles = func(filenames ...string) (*template.Template, error) {
			return nil, errors.New("error occurred")
		}
		homeHandler(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
