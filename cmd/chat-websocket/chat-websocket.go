package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rew150/chat-websocket/internal/router"
)

func main() {
	client := gin.Default()

	router.Route(client)

	client.Run()
}
