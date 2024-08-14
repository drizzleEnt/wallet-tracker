package handler

import (
	"github.com/drizzleent/wallet-tracker/backend/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.AuthService
}

func NewHandler(srv service.AuthService) *handler {
	return &handler{
		service: srv,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())

	api := router.Group("/")
	{
		api.POST("/register", h.Register)
		api.GET("/users/:address/nonce", h.UserNonce)
		api.POST("/signin", h.Signin)
		api.GET("/welcome", h.Welcome)
	}

	return router
}
