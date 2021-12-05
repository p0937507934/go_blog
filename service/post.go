package service

import (
	"github.com/blog/dto"
	"github.com/blog/model"
	"github.com/blog/repository"
)

type IPostService interface {
	Create(post *dto.CreatePostDto) error
	Update(post *dto.UpdatePostDto) error
	SelectById(id int) (*model.Post, error)
	Select() ([]*model.Post, error)
	DeleteByID(id int) error
	AllowUpdate(uid int, pid int) bool
}

type PostService struct {
	PostRepository repository.IPostRepository
}

func NewPostService(repo repository.IPostRepository) IPostService {
	return &PostService{repo}
}

func (p *PostService) Create(post *dto.CreatePostDto) error {
	err := p.PostRepository.Create(post)
	return err
}

func (p *PostService) Update(post *dto.UpdatePostDto) error {
	err := p.PostRepository.Update(post)
	return err
}

func (p *PostService) SelectById(id int) (*model.Post, error) {
	post, err := p.PostRepository.SelectById(id)
	return post, err
}

func (p *PostService) DeleteByID(id int) error {
	err := p.PostRepository.DeleteByID(id)
	return err
}

func (p *PostService) Select() ([]*model.Post, error) {
	postList := []*model.Post{}
	postList, err := p.PostRepository.Select(postList)
	return postList, err
}

func (p *PostService) AllowUpdate(uid int, pid int) bool {
	post, err := p.PostRepository.SelectById(pid)
	if err != nil || post.UID != uid {
		return false
	}
	return true
}
