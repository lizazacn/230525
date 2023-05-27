package API

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

// Options 跨域处理
func Options(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Headers", "*")
	context.Header("Access-Control-Allow-Credentials", "true")
	if context.Request.Method == http.MethodOptions {
		context.Status(201)
		context.Abort()
		return
	}
	context.Next()
}

func SocketIO() *socketio.Server {
	server := socketio.NewServer(nil)
	server.OnConnect("/", func(conn socketio.Conn) error {
		conn.Join("admin")
		return nil
	})
	// 设置用户在线状态
	server.OnEvent("/", "status", func(s socketio.Conn) {
		s.Emit("status", true)
	})
	// 管理端接收发送的消息
	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		s.Emit("msg", msg)
	})
	// 输入下一个字段
	server.OnEvent("/", "next", func(s socketio.Conn, msg string) {

		s.Emit("next", "下一个空间信息")
	})
	// 退出
	server.OnEvent("/", "bye", func(s socketio.Conn) {
		s.Emit("status", false)
		log.Printf("Exit SocketIO!")
		last := s.Context().(string)
		s.Emit("bye", last)
		err := s.Close()
		if err != nil {
			log.Printf("Close ScoketIO Error: %s", err)
			return
		}
		return
	})

	// 异常处理
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("error:", e)
	})

	// 断开连接
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	return server
}

func SaveData(ctx *gin.Context) {

}
