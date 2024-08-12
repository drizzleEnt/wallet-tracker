package api

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(c *gin.Context)
	UserNonce(c *gin.Context)
	Singin(c *gin.Context)
	Welcome(c *gin.Context)
	InitRoutes() *gin.Engine
}
