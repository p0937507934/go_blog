package repository

import (
	"github.com/blog/dto"
	"github.com/blog/model"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	Create(comment *dto.CreateCommentDto) error
	Update(comment *dto.UpdateCommentDto) error
	SelectById(id int) (*model.Comment, error)
	Select(comment []*model.Comment) ([]*model.Comment, error)
	DeleteByID(id int) error
}

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &CommentRepository{db}
}

func (c *CommentRepository) Create(comment *dto.CreateCommentDto) error {
	err := c.DB.Table("comment").Create(&comment).Error
	return err
}

func (c *CommentRepository) Update(comment *dto.UpdateCommentDto) error {
	err := c.DB.Table("comment").Where("id = ?", comment.ID).Save(&comment).Error
	return err
}

func (c *CommentRepository) SelectById(id int) (*model.Comment, error) {
	comment := &model.Comment{}
	err := c.DB.Preload("Post").Preload("User").First(&comment, id).Error
	return comment, err
}

func (c *CommentRepository) Select(post []*model.Comment) ([]*model.Comment, error) {
	commentList := []*model.Comment{}
	err := c.DB.Joins("Post").Joins("User").Find(&commentList).Error
	return commentList, err
}

func (c *CommentRepository) DeleteByID(id int) error {
	err := c.DB.Delete(&model.Comment{}, id).Error
	return err
}
