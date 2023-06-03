package user

import (
	"context"

	"github.com/gin-gonic/gin"
	userService "github.com/satanaroom/my-health/internal/service/user"
)

type Handler struct {
	ctx     context.Context
	service userService.Service
}

func NewHandler(ctx context.Context, service userService.Service) *Handler {
	return &Handler{
		ctx:     ctx,
		service: service,
	}
}

// POST		http://localhost:8080/user			- создать пользователя
// GET 		http://localhost:8080/user			- получить пользователей 			get
// GET 		http://localhost:8080/user/{id}		- получить пользователя по айди		get-by-id
// DELETE 	http://localhost:8080/user/delete	- удалить пользователя				delete
// PATCH	http://localhost:8080/user/update	- обновить пользователя				update

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	user := router.Group("/user")
	{
		user.POST("/", h.create)
		user.GET("/", h.get)
	}

	return router
}
