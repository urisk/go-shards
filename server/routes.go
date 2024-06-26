package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-shards/controllers"
	"go-shards/utils"
)

// NewRouter <function>
// is used to create a GIN engine instance where all controller and routes will be placed
func NewRouter() *gin.Engine {
	var envVars = utils.GetEnvVars()

	if envVars.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// middlewares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())

	// static files serving
	router.Static("/images", "./images")

	// endpoints
	v1 := router.Group("v1")
	{
		news := v1.Group("news")
		{
			nc := controllers.NewsController{}

			news.GET("/", nc.Get)
			news.GET("/sources", nc.GetSources)
			news.GET("/types", nc.GetTypes)
		}
	}

	return router
}
