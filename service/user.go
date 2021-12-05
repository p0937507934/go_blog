package service

import (
	"github.com/blog/dto"
	"github.com/blog/model"
	"github.com/blog/repository"
)

type IUserService interface {
	Create(user *dto.CreateUserDto) error
	// Update(user *model.User) error
	SelectById(id int) (model.User, error)
	Select(user *model.User) ([]model.User, error)
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo}
}

func (u *UserService) Create(user *dto.CreateUserDto) error {
	err := u.UserRepository.Create(user)
	return err
}

func (u *UserService) SelectById(id int) (model.User, error) {
	user, err := u.UserRepository.SelectById(id)
	return user, err
}

func (u *UserService) Select(user *model.User) ([]model.User, error) {
	findUser, err := u.UserRepository.Select(user)
	return findUser, err
}
