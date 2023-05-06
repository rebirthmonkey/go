package example

import (
	"testing"

	gosqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	db, mock, err := gosqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	rows := gosqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "Alice", "alice@example.com").
		AddRow(2, "Bob", "bob@example.com")

	mock.ExpectQuery("SELECT id, name, email FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	user, err := GetUser(db, 1)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	expectedUser := &User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	assert.Equal(t, expectedUser, user)
}
