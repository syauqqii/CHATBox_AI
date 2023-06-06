package router

import (
	"server/internal/ws"
	"server/internal/user"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/create_room", wsHandler.CreateRoom)
	r.GET("/ws/join_room/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/get_rooms", wsHandler.GetRooms)
	r.GET("ws/get_clients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
