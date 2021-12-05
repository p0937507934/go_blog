package repository

import (
	"github.com/blog/dto"
	"github.com/blog/model"
	"gorm.io/gorm"
)

type IPostRepository interface {
	Create(post *dto.CreatePostDto) error
	Update(post *dto.UpdatePostDto) error
	SelectById(id int) (*model.Post, error)
	Select(post []*model.Post) ([]*model.Post, error)
	DeleteByID(id int) error
}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &PostRepository{db}
}

func (p *PostRepository) Create(post *dto.CreatePostDto) error {
	err := p.DB.Table("post").Create(&post).Error
	return err
}

func (p *PostRepository) Update(post *dto.UpdatePostDto) error {
	err := p.DB.Table("post").Where("id = ?", post.ID).Save(&post).Error
	return err
}

func (p *PostRepository) SelectById(id int) (*model.Post, error) {
	post := &model.Post{}
	err := p.DB.Preload("User").Preload("Comments").First(&post, id).Error
	return post, err
}

func (p *PostRepository) Select(post []*model.Post) ([]*model.Post, error) {
	postList := []*model.Post{}
	err := p.DB.Preload("User").Preload("Comments").Find(&postList).Error
	return postList, err
}

func (p *PostRepository) DeleteByID(id int) error {
	err := p.DB.Delete(&model.Post{}, id).Error
	return err
}
