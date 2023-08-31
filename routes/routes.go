package routes

import (
	"github.com/gin-gonic/gin"

	controllers "gin_session_auth/controllers"
)

func HTMLRoutes(g *gin.RouterGroup) {
	g.GET("/policy", controllers.PolicyHandler())
	g.GET("/info", controllers.InfoHandler())
}

func RESTRoutes(g *gin.RouterGroup) {
	g.GET("/nodeInfo", controllers.NodeInfoHandler())
	g.GET("/podInfo", controllers.PodInfoHandler())
}
