package repositories

import (
	"database/sql"
	"goBase/app/schema"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
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

func (r *UserRepository) CheckUsername(UserName string) (bool, error) {
	query := `
		SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)
	`

	var exists bool
	err := r.DB.QueryRow(query, UserName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *UserRepository) GetUserByName(UserName string) (schema.UserSchema, error) {
	query := `
		SELECT id, username, email, firstname, lastname, is_superuser, password
		FROM users
		WHERE username = $1
	`

	var user schema.UserSchema
	err := r.DB.QueryRow(query, UserName).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.IsSuperuser,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}

	return user, nil
}
