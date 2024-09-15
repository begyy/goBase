package services

import (
	"errors"
	"fmt"
	"goBase/app/repositories"
	"goBase/app/schema"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{Repo: userRepository}
}

func (s *UserService) PasswordToHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s *UserService) SignUp(sch *schema.SignUpSchemaIn) (string, error) {
	checkUser, err := s.Repo.CheckUsername(sch.Username)
	if err != nil {
		return "", err
	}
	if checkUser {
		return "Username already exists", errors.New("username already exists")
	}

	password, err := s.PasswordToHashPassword(sch.Password)
	if err != nil {
		return "", err
	}

	_, err = s.Repo.CreateUser(
		sch.Username,
		sch.Email,
		sch.FirstName,
		sch.LastName,
		password,
		false)
	if err != nil {
		return "Can't create user", err
	}

	return "User signed up successfully", nil
}

func (s *UserService) SignIn(sch *schema.SignInSchemaIn) (schema.UserMeSchema, error) {
	getUser, err := s.Repo.GetUserByName(sch.Username)
	if err != nil {
		return getUser, err
	}
	passwordToHash, err := s.PasswordToHashPassword(sch.Password)
	if err != nil {
		return getUser, err
	}
	if CheckPasswordHash(sch.Password, passwordToHash) {
		return getUser, nil
	}
	return getUser, errors.New(fmt.Sprintf("username or password is incorrect"))

}
