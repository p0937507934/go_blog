package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blog/dto"
	"github.com/blog/model"
	"github.com/blog/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	// GetAll(c *gin.Context)
	// GetByID(c *gin.Context)
	// Delete(c *gin.Context)
	Create(c *gin.Context)
	Login(c *gin.Context)
}

type UserController struct {
	UserService service.IUserService
}

func NewUserController(service service.IUserService) IUserController {
	return &UserController{service}
}

// func (u *UserController) GetById(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Query("id"))
// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": "parameter error.",
// 		})
// 		return
// 	}
// 	user, err := u.UserService.SelectById(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": "user does not exist.",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": user,
// 	})
// }

func (u *UserController) Create(c *gin.Context) {

	user := &dto.CreateUserDto{}
	err := c.ShouldBindJSON(user)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Parameter Error.",
		})
		return
	}
	err = u.UserService.Create(user)
	if err != nil {
		fmt.Printf("%s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "create user succcess.",
	})

}

func (u *UserController) Login(c *gin.Context) {
	user := new(model.User)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("%s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	userList, err := u.UserService.Select(user)
	if err != nil {
		fmt.Printf("%s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if len(userList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "user not exists.",
		})
		return
	} else {
		//login and store to session using json.
		session := sessions.Default(c)
		userinfoSlice, _ := json.Marshal(userList)
		session.Set("userinfo", string(userinfoSlice))
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"msg": "login success.",
		})
	}

}

// func (u *UserController) Update(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Query("id"))
// 	updatedUser := new(model.User)
// 	if err != nil {
// 		fmt.Printf("%s", err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": "parameter error.",
// 		})
// 		return
// 	}
// 	err = c.ShouldBindJSON(&updatedUser)
// 	if err != nil {
// 		fmt.Printf("%s", err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": err.Error(),
// 		})
// 		return
// 	}
// 	user, err := u.UserService.SelectById(id)
// 	if err != nil {
// 		fmt.Printf("%s", err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": "user not found.",
// 		})
// 		return
// 	}

// }
