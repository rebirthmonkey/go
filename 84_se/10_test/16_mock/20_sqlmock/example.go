package example

import (
	"database/sql"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func GetUser(db *sql.DB, id int) (*User, error) {
	user := &User{}
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
