package main

import (
	"fmt"
	userHandler "github.com/satanaroom/my-health/internal/handler/user"
	userServer "github.com/satanaroom/my-health/internal/server/user"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := new(userServer.Server)

	handler := userHandler.NewHandler()

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
