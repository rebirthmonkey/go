package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	expectedUser := &User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users/1", r.URL.Path)

		resp := &User{
			ID:    1,
			Name:  "Alice",
			Email: "alice@example.com",
		}
		body, _ := json.Marshal(resp)

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	userService := &UserService{
		baseURL: server.URL,
	}

	user, err := userService.GetUserByID(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}
