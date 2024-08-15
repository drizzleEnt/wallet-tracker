package handler

import (
	"github.com/drizzleent/wallet-tracker/backend/internal/middleware"
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

	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AddAllowHeaders("user")
	cfg.AddAllowHeaders("authorization")

	router.Use(cors.New(cfg))

	api := router.Group("/")
	{
		api.POST("/register", h.Register)
		api.GET("/users/:address/nonce", h.UserNonce)
		api.POST("/signin", h.Signin)
		api.GET("/welcome", middleware.AuthMiddleware(h.service), h.Welcome)
	}

	return router
}
