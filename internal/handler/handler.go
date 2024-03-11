package handlers

import (
	"github.com/Waldemarsch/medods_test/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	tokens := router.Group("/token")
	{
		tokens.POST("/create", h.CreateToken)
		tokens.PUT("/refresh", h.RefreshToken)
	}

	return router
}
