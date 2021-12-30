package routers

import (
	"gin-blog/pkg/setting"
	v1 "gin-blog/routers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/articles/:id", v1.GetArticle)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("/articles/:id", v1.EditArticle)
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
