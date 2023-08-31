package controllers

import (
	"gin_session_auth/helpers"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PolicyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "policy.html", nil)
	}
}

func InfoHandler() gin.HandlerFunc {
	infos := helpers.GetClusterInfo()
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "information.html", gin.H{
			"clusterNode_GPU": template.HTML(infos),
		})
	}
}

func HomeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	}
}

type NodeInfoHTML struct {
	InnerHTML string `json:"innerHTML"`
}

func NodeInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramMap := c.Request.URL.Query()

		clusterID := paramMap["clusterNum"][0]
		nodeCount := paramMap["nodeCnt"][0]
		gpuCount := paramMap["gpuCnt"][0]
		nHtml := &NodeInfoHTML{}
		nHtml.InnerHTML = helpers.GetNodeInfo(clusterID, nodeCount, gpuCount)

		c.JSON(http.StatusOK, nHtml)
	}
}

type PodInfoHTML struct {
	InnerHTML string `json:"innerHTML"`
}

func PodInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramMap := c.Request.URL.Query()

		clusterID := paramMap["clusterNum"][0]
		pHtml := &PodInfoHTML{}
		pHtml.InnerHTML = helpers.GetPodInfo(clusterID)

		c.JSON(http.StatusOK, pHtml)
	}
}
