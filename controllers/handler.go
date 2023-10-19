package controllers

import (
	"gin_session_auth/helpers"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

func PolicyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "policy.html", nil)
	}
}

func InfoHandler() gin.HandlerFunc {
	infos := helpers.GetClusterList()
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
		node := c.DefaultQuery("node", "")
		nHtml := &NodeInfoHTML{}
		nHtml.InnerHTML = helpers.GetNodeMetricInfo(node, helpers.GetNodeInfo(node))

		c.JSON(http.StatusOK, nHtml)
	}
}

type PodInfoHTML struct {
	PodHTML string `json:"podHTML"`
}

func PodInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		node := c.DefaultQuery("node", "")
		klog.Infoln(node)
		pHtml := &PodInfoHTML{}
		pHtml.PodHTML = helpers.GetPodInfo(node)

		c.JSON(http.StatusOK, pHtml)
	}
}

type ClusterInfoHTML struct {
	InnerHTML string `json:"innerHTML"`
}

func ConfirmChanges() gin.HandlerFunc {
	return func(c *gin.Context) {
		cluster := c.DefaultQuery("cluster", "")
		nHtml := &ClusterInfoHTML{}
		nHtml.InnerHTML = helpers.GetClusterInfo(cluster)
		c.JSON(http.StatusOK, nHtml)
	}
}
