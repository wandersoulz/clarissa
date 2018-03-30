package web

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	cors "github.com/itsjamie/gin-cors"
)

var socketio_Server *socketio.Server

// ConvertTime - convert simulation time to real time
func ConvertTime(time_seconds float64) float64 {
	milliseconds := time_seconds * 100
	return milliseconds
}

func getElapsedTime(prevTime float64, currentTime float64) float64 {
	return currentTime - prevTime
}

func ConvertElapsedTimeToRealTime(elapsed float64) float64 {
	convCoef := 60.0
	return elapsed / convCoef
}

func main() {
	r := gin.Default()
	var err error

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: true,
	}))

	socketio_Server, err = socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	r.GET("/socket.io", socketHandler)
	r.POST("/socket.io", socketHandler)
	r.Handle("WS", "/socket.io", socketHandler)
	r.Handle("WSS", "/socket.io", socketHandler)

	r.Group("/socket/")
}

func socketHandler(c *gin.Context) {
	socketio_Server.On("connection", func(so socketio.Socket) {
		fmt.Println("on connection")

		so.Join("chat")

		so.On("chat message", func(msg string) {
			fmt.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			fmt.Println("on disconnect")
		})
	})

	socketio_Server.On("error", func(so socketio.Socket, err error) {
		fmt.Printf("[ WebSocket ] Error : %v", err.Error())
	})

	socketio_Server.ServeHTTP(c.Writer, c.Request)
}
