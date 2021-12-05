package service

import (
	"errors"

	"github.com/blog/dto"
	"github.com/blog/model"
	"github.com/blog/repository"
)

type ICommentService interface {
	Create(comment *dto.CreateCommentDto) error
	Update(comment *dto.UpdateCommentDto) error
	SelectById(id int) (*model.Comment, error)
	Select() ([]*model.Comment, error)
	DeleteByID(id int) error
	AllowUpdate(uid int, cid int) bool
}

type CommentService struct {
	CommentRepository repository.ICommentRepository
	PostRepository    repository.IPostRepository
}

func NewCommentService(commentRepo repository.ICommentRepository, postRepo repository.IPostRepository) ICommentService {
	return &CommentService{commentRepo, postRepo}
}

func (c *CommentService) Create(comment *dto.CreateCommentDto) error {
	post, err := c.PostRepository.SelectById(comment.PID)
	if err != nil {
		return err
	}
	if post.Status == 1 {
		return errors.New("This post can not be comment.")
	}
	err = c.CommentRepository.Create(comment)
	return err
}

func (c *CommentService) Update(comment *dto.UpdateCommentDto) error {
	err := c.CommentRepository.Update(comment)
	return err
}

func (c *CommentService) SelectById(id int) (*model.Comment, error) {
	comment, err := c.CommentRepository.SelectById(id)
	return comment, err
}

func (c *CommentService) DeleteByID(id int) error {
	err := c.CommentRepository.DeleteByID(id)
	return err
}

func (c *CommentService) Select() ([]*model.Comment, error) {
	commentList := []*model.Comment{}
	commentList, err := c.CommentRepository.Select(commentList)
	return commentList, err
}
func (c *CommentService) AllowUpdate(uid int, cid int) bool {
	comment, err := c.CommentRepository.SelectById(cid)
	if err != nil || comment.UID != uid {
		return false
	}
	return true
}
