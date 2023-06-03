package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	userHandler "github.com/satanaroom/my-health/internal/handler/user"
	userServer "github.com/satanaroom/my-health/internal/server/user"
	userService "github.com/satanaroom/my-health/internal/service/user"
)

func main() {
	ctx := context.Background()

	srv := new(userServer.Server)

	service := userService.NewService()
	handler := userHandler.NewHandler(ctx, service)

	fmt.Println("my-health app started")
	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			log.Fatalf("run server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	fmt.Println("my-health app stopped")
}
