package user

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) get(c *gin.Context) {
	c.JSON(200, "got users")
}
