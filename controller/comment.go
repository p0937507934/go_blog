package controller

import (
	"net/http"
	"strconv"

	"github.com/blog/dto"
	"github.com/blog/service"
	"github.com/blog/utils"
	"github.com/gin-gonic/gin"
)

type ICommentController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
}

type CommentController struct {
	CommentService service.ICommentService
}

func NewCommentController(service service.ICommentService) ICommentController {
	return &CommentController{service}
}

func (cm *CommentController) GetAll(c *gin.Context) {
	commentList, err := cm.CommentService.Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Get Data Failed.",
		})
		return
	}
	c.JSON(http.StatusOK, commentList)

}

func (cm *CommentController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}

	comment, err := cm.CommentService.SelectById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Get Data Failed.",
		})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (cm *CommentController) Update(c *gin.Context) {
	var role int16
	var id int
	cid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	comment := new(dto.UpdateCommentDto)
	comment.ID = cid
	err = c.ShouldBind(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	role = utils.GetSessionRole(c)
	id = utils.GetSessionID(c)
	if role == 0 {
		ok := cm.CommentService.AllowUpdate(id, comment.ID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Auth Failed.",
			})
			return
		}
	}
	err = cm.CommentService.Update(comment)
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

func (cm *CommentController) Delete(c *gin.Context) {
	cid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	id := utils.GetSessionID(c)
	role := utils.GetSessionRole(c)
	if role == 0 {
		ok := cm.CommentService.AllowUpdate(id, cid)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Auth Failed.",
			})
			return
		}
	}
	err = cm.CommentService.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Delete Failed.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete success.",
	})

}

func (cm *CommentController) Create(c *gin.Context) {
	var comment *dto.CreateCommentDto
	err := c.ShouldBind(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	id := utils.GetSessionID(c)
	comment.UID = id
	err = cm.CommentService.Create(comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Create Comment Failed.",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "Create Comment Successed.",
	})
}
