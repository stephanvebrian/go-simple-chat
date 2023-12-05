package handler

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestNewHome(t *testing.T) {
	router := mux.NewRouter()

	resultRouter := NewHome(router)

	assert.Equal(t, router, resultRouter, "Router mismatch")
}
