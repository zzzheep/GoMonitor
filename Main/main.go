package main

import (
	"GoMonitor/Model"
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
	ticker := time.NewTicker(time.Second * 2)
	for range ticker.C {
		processInfo := Model.GetProcessInfo()
		err := conn.WriteJSON(processInfo)
		if err != nil {
			fmt.Println("senderr", err)
		}
	}

	// // 必须死循环，gin通过协程调用该handler函数，一旦退出函数，ws会被主动销毁
	// for {
	// 	// recieve
	// 	_, reply, err := conn.ReadMessage()
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Println(string(reply))
	// }
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.LoadHTMLGlob("../Views/*")
	r.GET("/monitor", func(c *gin.Context) {
		WsHandler(c.Writer, c.Request)
	})
	r.Run("localhost:12312")
}
