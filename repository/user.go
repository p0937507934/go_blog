package repository

import (
	"fmt"

	"github.com/blog/dto"
	"github.com/blog/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *dto.CreateUserDto) error
	Update(user model.User) error
	SelectById(id int) (model.User, error)
	Select(user *model.User) ([]model.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Create(user *dto.CreateUserDto) error {
	err := u.DB.Table("user").Create(&user).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func (u *UserRepository) Update(user model.User) error {
	err := u.DB.Save(&user).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func (u *UserRepository) SelectById(id int) (model.User, error) {
	user := model.User{}
	err := u.DB.First(&user, id).Error
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, err
}

func (u *UserRepository) Select(user *model.User) ([]model.User, error) {
	userList := []model.User{}
	result := u.DB.Where(&user).Find(&userList)
	if result.Error != nil {
		fmt.Println(result.Error)
		return userList, result.Error
	}

	return userList, result.Error
}
