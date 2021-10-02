package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/rew150/chat-websocket/internal/router"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	socketSrv := socketio.NewServer(nil)

	router.SocketIoRoute(socketSrv)

	ginHandler := gin.Default()

	router.Route(ginHandler)
	ginHandler.GET("/socket.io/*any", gin.WrapH(socketSrv))
	ginHandler.POST("/socket.io/*any", gin.WrapH(socketSrv))
	ginHandler.StaticFS("/public", http.Dir("public"))

	ginSrv := &http.Server{
		Addr:    ":8080",
		Handler: ginHandler,
	}

	go func() {
		if err := socketSrv.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	go func() {
		if err := ginSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http listen error: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := socketSrv.Close(); err != nil {
		log.Fatalf("Server shutdown failed: %s \n", err)
	}
	if err := ginSrv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s \n", err)
	}

	log.Println("Exit")
}
