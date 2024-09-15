package repositories

import (
	"database/sql"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(
	UserName, Email, FirstName, LastName, Password string, isSuperUser bool,
) (int, error) {
	query := `
		INSERT INTO users (username, email, firstname, lastname, password, is_superuser)
		VALUES ($1, $2, $3, $4,$5, $6)
		RETURNING id;
	`
	var userID int
	err := r.DB.QueryRow(query, UserName, Email, FirstName, LastName, Password, isSuperUser).Scan(&userID)
	if err != nil {
		log.Fatal("Failed to insert user:", err)
		return -1, err
	}
	return userID, nil
}
