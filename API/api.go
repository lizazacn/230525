package API

import (
	"230525/Struct"
	"encoding/json"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"sync"
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
	var ConnMap sync.Map
	server := socketio.NewServer(nil)
	// 客户端连接，不加入房间，将连接信息存入连接map种
	server.OnConnect("/client", func(conn socketio.Conn) error {
		conn.Emit("/connect", conn.ID())
		return nil
	})
	// 管理端连接，全部加入admin房间
	server.OnConnect("/admin", func(conn socketio.Conn) error {
		conn.Join("admin")
		return nil
	})
	// 管理员端更新所有用户的状态
	server.OnEvent("/client", "status", func(conn socketio.Conn, msg string) {
		// 将所有连接存入ConnMap
		ConnMap.Store(conn.ID(), conn)
		server.BroadcastToRoom("/admin", "admin", "status", msg)
	})
	// 管理端接收发送的消息
	server.OnEvent("/client", "msg", func(s socketio.Conn, msg string) {
		server.BroadcastToRoom("/admin", "admin", "status", msg)
	})
	// 让指定用户输入下一个字段
	server.OnEvent("/admin", "next", func(s socketio.Conn, msg string) {
		var m = new(Struct.Msg)
		err := json.Unmarshal([]byte(msg), m)
		if err != nil {
			return
		}
		m.Status = false
		marshal, _ := json.Marshal(m)
		c, ok := ConnMap.Load(m.ID)
		if !ok || c == nil {
			server.BroadcastToRoom("/admin", "admin", "status", string(marshal))
			return
		}
		var conn = c.(socketio.Conn)
		conn.Emit("/next", string(marshal))
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
	server.OnError("/client", func(s socketio.Conn, e error) {
		log.Println("error:", e)
	})

	// 断开连接
	server.OnDisconnect("/client", func(s socketio.Conn, reason string) {
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
