package user

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// POST http://localhost:8080/user/create - создать пользователя
// GET http://localhost:8080/user/get - создать пользователя

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	user := router.Group("/user")
	{
		user.POST("/create", h.create)
		user.GET("/get", h.get)
	}

	return router
}
