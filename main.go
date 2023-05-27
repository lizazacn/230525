package main

import (
	"230525/API"
	"230525/Service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := Service.InitSQLConn()
	if err != nil {
		log.Printf("连接数据库异常：%v", err)
	}
	r := gin.Default()
	r.Use(API.Options)
	server := API.SocketIO()
	r.Any("/socket.io/*any", gin.WrapH(server))
	r.POST("/saveData", API.SaveData)
	err = r.Run("0.0.0.0:80")
	if err != nil {
		return
	}
}
