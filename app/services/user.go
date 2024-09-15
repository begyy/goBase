package services

import (
	"goBase/app/repositories"
	"goBase/app/schema"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{Repo: userRepository}
}

func (s *UserService) SignUp(sch schema.SignUpSchemaIn) error {
	//s.Repo.CreateUser(sch.Username, sch.Email, sch.FirstName, sch.LastName, sch.Password)
	//return nil
	return nil
}
