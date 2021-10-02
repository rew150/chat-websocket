package handler

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/rew150/chat-websocket/internal/socketctx"
)

func OnSocketIoConnectHandler(s socketio.Conn) error {
	s.SetContext(socketctx.BlankCtx)
	log.Println("Socket/connected:", s.ID())
	return nil
}

func OnSocketIoDisconnectHandler(s socketio.Conn, reason string) {
	log.Println("Socket/disconnected:", s.ID(), reason)
}

func OnSocketIoErrorHandler(s socketio.Conn, err error) {
	log.Println("Socket/error:", s.ID(), err)
}

func OnSocketIoEventMessageHandler(s socketio.Conn, msg string) {
	log.Println("Socket/event/message:", s.ID(), s.Context())
	s.Emit("message", msg)
}
