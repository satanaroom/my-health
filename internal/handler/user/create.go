package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) create(c *gin.Context) {
	if err := h.service.Create(h.ctx); err != nil {
		c.JSON(500, fmt.Sprintf("create user error: %s", err.Error()))
	}
	c.JSON(200, "user created")
}
