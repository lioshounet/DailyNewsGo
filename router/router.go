package router

import (
	"github.com/gin-gonic/gin"
	"thor/router/daily"
	"thor/router/user"
)

func RegisterRoute(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRouter(router)
	dailyRouter(router)
}

func userRouter(router *gin.Engine) {
	rp := router.Group("/user")
	{
		rp.POST("/login", user.Login)
		rp.POST("/userinfo", user.UserInfo)
		rp.POST("/edituserbyuid", user.EditUserByUid)
	}

	rg := router.Group("/user")
	{
		rg.GET("/login", user.Login)
		rg.GET("/userinfo", user.UserInfo)
		rg.GET("/edituserbyuid", user.EditUserByUid)
	}
}

func dailyRouter(router *gin.Engine) {
	rp := router.Group("/daily")
	{
		rp.POST("/list", daily.List)
		rp.POST("/create", daily.Create)
		rp.POST("/listbydate", daily.ListByDate)
		rp.POST("/detailbyid", daily.DetailById)
	}

	rg := router.Group("/daily")
	{
		rp.GET("/list", daily.List)
		rg.GET("/create", daily.Create)
		rg.GET("/listbydate", daily.ListByDate)
		rp.GET("/detailbyid", daily.DetailById)
	}
}
