package handler

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stphanvebrian/go-simple-chat/usecase/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewChat(t *testing.T) {
	router := mux.NewRouter()
	mockDB := &repository.MockDB{}

	resultRouter := NewChat(router, mockDB)
	assert.Equal(t, router, resultRouter, "Router mismatch")
}

func TestSaveMessage(t *testing.T) {
	mockDB := &repository.MockDB{}
	mockResult := &repository.MockResult{}
	mockDB.On("Exec", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResult, nil)

	err := saveMessage(mockDB, "sender", "recipient", "Test Message")

	assert.NoError(t, err, "Error in SaveMessage function")
	mockDB.AssertCalled(t, "Exec", "INSERT INTO messages (sender, recipient, message) VALUES (?, ?, ?)", "sender", "recipient", "Test Message")
}
