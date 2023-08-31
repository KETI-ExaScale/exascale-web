package main

import (
	"github.com/gin-gonic/gin"

	//"html/template"
	//"strings"

	globals "gin_session_auth/globals"
	routes "gin_session_auth/routes"
)

func main() {
	router := gin.Default()

	globals.InitClient()

	router.Static("/assets", "./assets")
	router.Static("/static/css", "./static/css")
	router.Static("/static/img", "./static/img")
	router.Static("/static/scss", "./static/scss")
	router.Static("/static/vendor", "./static/vendor")
	router.Static("/static/js", "./static/js")

	router.LoadHTMLGlob("templates/*.html")
	public := router.Group("/")
	routes.HTMLRoutes(public)
	routes.RESTRoutes(public)

	router.Run("0.0.0.0:8080")
}
