package controllers

import (
	"fmt"
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
	manager := helpers.NewNodeManager()
	infos := manager.GetClusterList()
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "information.html", gin.H{
			"clusterRadio": template.HTML(infos),
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
		// node := c.DefaultQuery("node", "")
		nHtml := &NodeInfoHTML{}
		// manager := helpers.NewNodeManager()
		// nHtml.InnerHTML = manager.GetNodeMetricInfo(node, manager.GetNodeGPUInfo(node))

		c.JSON(http.StatusOK, nHtml)
	}
}

type PodInfoHTML struct {
	PodHTML string `json:"podHTML"`
}

func PodInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// node := c.DefaultQuery("node", "")
		// klog.Infoln(node)
		pHtml := &PodInfoHTML{}
		// manager := helpers.NewNodeManager()
		// pHtml.PodHTML = manager.GetPodInfo(node)

		c.JSON(http.StatusOK, pHtml)
	}
}

type ClusterInfoHTML struct {
	InnerHTML string `json:"innerHTML"`
}

func ConfirmChanges() gin.HandlerFunc {
	fmt.Println("ConfirmChanges")
	return func(c *gin.Context) {
		nHtml := &ClusterInfoHTML{}
		manager := helpers.NewNodeManager()
		nHtml.InnerHTML = manager.GetClusterInfo("Cluster1")
		c.JSON(http.StatusOK, nHtml)
	}
}
