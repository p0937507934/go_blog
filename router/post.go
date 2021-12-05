package router

import (
	"github.com/blog/controller"
	"github.com/blog/db"
	"github.com/blog/middleware"
	"github.com/blog/repository"
	"github.com/blog/service"
	"github.com/gin-gonic/gin"
)

func PostRouter(r *gin.Engine) {

	PostRepository := repository.NewPostRepository(db.DB)
	PostService := service.NewPostService(PostRepository)
	PostController := controller.NewPostController(PostService)

	r.POST("/post", middleware.Auth, PostController.Create)
	r.GET("/post", middleware.Auth, PostController.GetAll)
	r.GET("/post/:id", middleware.Auth, PostController.GetByID)
	r.PUT("/post/:id", middleware.Auth, PostController.Update)
	r.PUT("/post/lock/:id", middleware.AuthAdmin, PostController.LockPost)
	r.DELETE("/post/:id", middleware.Auth, PostController.Delete)
}
