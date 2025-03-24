package service

import "testDBMock/internal/repository"

type UserService interface {
	GetUser(id int) (*repository.User, error)
	CreateUser(username, email string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUser(id int) (*repository.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) CreateUser(username, email string) error {
	user := &repository.User{
		Username: username,
		Email:    email,
	}
	return s.userRepo.Create(user)
}
