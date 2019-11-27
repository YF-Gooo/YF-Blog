package server

import (
	"singo/api"
	"singo/middleware"
	"os"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())
	r.Static("/static", "/static")
	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// User Routing
			authed.GET("user/me", api.UserMe)
			authed.DELETE("user/logout", api.UserLogout)
		}
		v1.POST("uploadimage",api.UploadImage)
		v1.GET("listimage",api.ListImage) 
		v1.POST("article", api.CreateArticle)
		v1.GET("article/:id", api.ShowArticle)
		v1.GET("articles", api.ListArticle)
		v1.GET("userarticles", api.UserListArticle)
		v1.GET("searcharticles", api.SearchListArticle)
		v1.GET("usersearcharticles", api.UserSearchListArticle)
		v1.PUT("article/:id", api.UpdateArticle)
		v1.DELETE("article/:id", api.DeleteArticle)

		v1.POST("video", api.CreateVideo)
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		v1.PUT("video/:id", api.UpdateVideo)
		v1.DELETE("video/:id", api.DeleteVideo)
	}
	return r
}

