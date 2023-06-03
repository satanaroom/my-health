package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) create(c *gin.Context) {
	fmt.Println("hello, user created")

	c.JSON(200, "user created")
}
