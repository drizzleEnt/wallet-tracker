package handler

import "github.com/gin-gonic/gin"

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/")
	{
		api.POST("/register", h.Register)
		api.GET("/users/{address:^0x[a-fA-F0-9]{40}$}/nonce", h.UserNonce)
		api.POST("/signin", h.Signin)
		api.GET("/welcome", h.Welcome)
	}

	return router
}
