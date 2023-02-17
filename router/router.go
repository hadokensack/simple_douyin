package router

import (
	"simple_douyin/controller/user"
	"simple_douyin/controller/video"
	"simple_douyin/models"

	"github.com/gin-gonic/gin"
)

func Init_Router() *gin.Engine {
	models.Init_DB() //创建数据库
	r := gin.Default()

	r.Static("static", "./static")  //静态资源
	apiRouter := r.Group("/douyin") //注册douyin分组路由

	//basic apis
	apiRouter.GET("/feed/", video.FeedVideoListHandler)
	apiRouter.GET("/user/", user.UserInfoHandler)
	apiRouter.POST("/user/login/", user.UserLoginHandler)
	apiRouter.POST("/user/register/", user.UserRegisterHandler)
	apiRouter.POST("/publish/action/", video.PublishVideoHandler)
	apiRouter.GET("/publish/list/", video.QueryVideoListHandler)

	return r
}
