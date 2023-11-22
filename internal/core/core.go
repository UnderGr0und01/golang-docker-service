package core

import "github.com/gin-gonic/gin"

type Controller interface {
	GetContainers(c *gin.Context)
	StartContainer(c *gin.Context)
	StopContainer(c *gin.Context)
	GetLogs(c *gin.Context)
}
