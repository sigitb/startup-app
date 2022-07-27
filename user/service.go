package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
	Login(LoginInput LoginInput) (User, error)
}

type service struct {
	repository Respository
}

func NewService(repository Respository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.MinCost)
	if err != nil{
		return user, err
	}
	user.Password = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil{
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error){
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil{
		return user, err
	}
	if user.Id == 0{
		return user, errors.New("No user found the email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user , err
	}
	return user ,nil
}