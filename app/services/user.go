package services

import (
	"errors"
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
	passwordWithSecret := password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSecret), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *UserService) CheckPasswordHash(password, hashedPassword string) bool {
	passwordWithSecret := password

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordWithSecret))
	return err == nil
}

func (s *UserService) SignUp(sch *schema.SignUpSchemaIn) (string, error) {
	checkUser, err := s.Repo.CheckUsername(sch.Username)
	if err != nil {
		return "", err
	}
	if checkUser {
		return "", errors.New("username already exists")
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
		return "", err
	}

	return "User signed up successfully", nil
}

func (s *UserService) SignIn(sch *schema.SignInSchemaIn) (*schema.UserMeSchema, error) {
	user, err := s.Repo.GetUserByName(sch.Username)
	if err != nil {
		return nil, errors.New("username or password is incorrect")
	}
	if !s.CheckPasswordHash(sch.Password, user.Password) {

		return nil, errors.New("username or password is incorrect")
	}
	userMe := &schema.UserMeSchema{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		IsSuperuser: user.IsSuperuser,
	}
	return userMe, nil
}
