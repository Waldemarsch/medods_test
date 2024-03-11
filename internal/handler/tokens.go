package handlers

import (
	"github.com/Waldemarsch/medods_test/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateToken(c *gin.Context) {
	Token := new(models.Token)
	queryParams := c.Request.URL.Query()
	Token.CreateToken.Request.GUID = queryParams.Get("guid")
	h.services.Tokenization.CreateToken(c, Token)
	c.JSON(http.StatusOK, Token.CreateToken.Response)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	Token := new(models.Token)
	Token.RefreshToken.Request.Access = c.Query("access")
	Token.RefreshToken.Request.Refresh = c.Query("refresh")
	h.services.Tokenization.RefreshToken(c, Token)
	c.JSON(http.StatusOK, Token.RefreshToken.Response)
}
