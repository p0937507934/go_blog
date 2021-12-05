package router

import (
	"github.com/blog/controller"
	"github.com/blog/db"
	"github.com/blog/middleware"
	"github.com/blog/repository"
	"github.com/blog/service"
	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.Engine) {

	CommentRepository := repository.NewCommentRepository(db.DB)
	PostRepository := repository.NewPostRepository(db.DB)
	CommentService := service.NewCommentService(CommentRepository, PostRepository)
	CommentController := controller.NewCommentController(CommentService)

	r.POST("/comment", middleware.Auth, CommentController.Create)
	r.GET("/comment", middleware.Auth, CommentController.GetAll)
	r.GET("/comment/:id", middleware.Auth, CommentController.GetByID)
	r.PUT("/comment/:id", middleware.Auth, CommentController.Update)
	r.DELETE("/comment/:id", middleware.Auth, CommentController.Delete)
}
