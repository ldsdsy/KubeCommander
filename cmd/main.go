package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ldsdsy/KubeCommander/cluster"
)

func main() {
	r := gin.Default()

	r.POST("/clusters", cluster.AddCluster)
	r.DELETE("/clusters/:id", cluster.DeleteCluster)
	r.GET("/clusters", cluster.ListCluster)

	r.Run() // 默认在8080端口监听
}
