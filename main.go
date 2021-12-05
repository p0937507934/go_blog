package main

import (
	"github.com/blog/db"
	"github.com/blog/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitMysql()

	r := gin.Default()
	store := cookie.NewStore([]byte("hsipl"))
	r.Use(sessions.Sessions("mysession", store))
	router.UserRouter(r)
	router.PostRouter(r)
	router.CommentRouter(r)
	router.SocketRouter(r)
	r.Run()
}
