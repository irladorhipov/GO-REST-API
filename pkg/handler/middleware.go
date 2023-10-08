package handler

import (
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	autharizationHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(autharizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}