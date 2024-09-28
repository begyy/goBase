package repositories

import "database/sql"

type UserTokenRepository struct {
	DB *sql.DB
}

func NewUserTokenRepository(db *sql.DB) *UserTokenRepository {
	return &UserTokenRepository{DB: db}
}

func (r *UserTokenRepository) GetUserByToken(token string) (int, error) {
	query := `SELECT user_id FROM user_token WHERE token = $1`
	var userID int
	err := r.DB.QueryRow(query, token).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *UserTokenRepository) AddUserToken(userID int, token string) error {
	query := `INSERT INTO user_token (user_id, token) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, userID, token)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserTokenRepository) DeleteUserToken(userID int) error {
	query := `DELETE FROM user_token WHERE user_id = $1`
	_, err := r.DB.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserTokenRepository) GetUserTokenByUserId(userID int) (string, error) {
	query := `SELECT token FROM user_token WHERE user_id = $1`
	var token string
	err := r.DB.QueryRow(query, userID).Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}
