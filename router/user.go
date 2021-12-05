package router

import (
	"github.com/blog/controller"
	"github.com/blog/db"
	"github.com/blog/repository"
	"github.com/blog/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	UserRepository := repository.NewUserRepository(db.DB)
	UserService := service.NewUserService(UserRepository)
	UserController := controller.NewUserController(UserService)
	r.POST("/user/register", UserController.Create)
	// r.GET("/user", UserController.GetById)
	r.POST("/user/login", UserController.Login)

}
