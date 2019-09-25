package actions

import (
	"api-framework/actions/api"
	"api-framework/core/config"
	"github.com/gin-gonic/gin"
	_ "api-framework/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(config.ServerConfig.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.JWT());
	apiv1.Use()
	{
		apiv1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}

	return r
}
