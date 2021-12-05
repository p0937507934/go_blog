package controller

import (
	"net/http"
	"strconv"

	"github.com/blog/dto"
	"github.com/blog/service"
	"github.com/blog/utils"
	"github.com/gin-gonic/gin"
)

type IPostController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
	LockPost(c *gin.Context)
}

type PostController struct {
	PostService service.IPostService
}

func NewPostController(service service.IPostService) IPostController {
	return &PostController{service}
}

func (p *PostController) GetAll(c *gin.Context) {
	postList, err := p.PostService.Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Get Data Failed.",
		})
		return
	}
	c.JSON(http.StatusOK, postList)

}

func (p *PostController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}

	post, err := p.PostService.SelectById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Get Data Failed.",
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (p *PostController) Update(c *gin.Context) {
	var role int16
	var id int
	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	var post = new(dto.UpdatePostDto)
	post.ID = pid
	err = c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	role = utils.GetSessionRole(c)
	id = utils.GetSessionID(c)
	if role == 0 {
		ok := p.PostService.AllowUpdate(id, post.ID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Auth Failed.",
			})
			return
		}
	}
	err = p.PostService.Update(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Updated Failed.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "update success.",
	})
}

func (p *PostController) Delete(c *gin.Context) {
	var id int
	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	id = utils.GetSessionID(c)
	role := utils.GetSessionRole(c)
	if role == 0 {
		ok := p.PostService.AllowUpdate(id, pid)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Auth Failed.",
			})
			return
		}
	}
	err = p.PostService.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Deleted Failed.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete Success.",
	})

}

func (p *PostController) Create(c *gin.Context) {
	var post *dto.CreatePostDto
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	id := utils.GetSessionID(c)
	post.UID = id
	err = p.PostService.Create(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Create Post Failed.",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "Create Post Successed.",
	})

}

func (p *PostController) LockPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	post := new(dto.UpdatePostDto)
	post.ID = id
	post.Status = 1
	err = p.PostService.Update(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Lock Post Failed.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lock Post Success.",
	})
}
