package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理ws请求
func WsHandler(w http.ResponseWriter, r *http.Request) {
	var conn *websocket.Conn
	var err error
	conn, err = wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	// 必须死循环，gin通过协程调用该handler函数，一旦退出函数，ws会被主动销毁
	for {
		// recieve
		_, reply, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(reply))
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("../Views/*")
	r.GET("/monitor", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "monitorservice",
		})
	})
	r.GET("/monitor/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": name,
		})
	})
	r.Run()
}
