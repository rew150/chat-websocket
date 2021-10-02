package router

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/rew150/chat-websocket/internal/handler"
)

func SocketIoRoute(srv *socketio.Server) {
	srv.OnConnect("/", handler.OnSocketIoConnectHandler)
	srv.OnDisconnect("/", handler.OnSocketIoDisconnectHandler)
	srv.OnError("/", handler.OnSocketIoErrorHandler)

	srv.OnEvent("/", "message", handler.OnSocketIoEventMessageHandler)
}
