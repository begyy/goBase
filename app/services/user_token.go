package services

import (
	"crypto/rand"
	"encoding/base64"
	"goBase/app/repositories"
)

type UserTokenService struct {
	Repo *repositories.UserTokenRepository
}

func generateToken(n int) (string, error) {
	b := make([]byte, n*3/4)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func NewUserTokenService(userTokenRepository *repositories.UserTokenRepository) *UserTokenService {
	return &UserTokenService{Repo: userTokenRepository}
}

func (s *UserTokenService) GetUserByToken(token string) (int, error) {
	return s.Repo.GetUserByToken(token)
}

func (s *UserTokenService) GetTokenOrAddToken(userId int) (string, error) {
	getToken, err := s.Repo.GetUserTokenByUserId(userId)
	if err != nil {
		generateToken, err := generateToken(48)
		if err != nil {
			return "", err
		}
		err = s.Repo.AddUserToken(userId, generateToken)
		if err != nil {
			return "", err
		}
		return generateToken, nil
	}
	return getToken, nil
}
