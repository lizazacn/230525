package main

import (
	"230525/API"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(API.Options)
	server := API.SocketIO()
	r.Any("/socket.io/*any", gin.WrapH(server))
	r.POST("/saveData", API.SaveData)
	err := r.Run("0.0.0.0:80")
	if err != nil {
		return
	}
}
