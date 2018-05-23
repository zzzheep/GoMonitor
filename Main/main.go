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

var connMap = make(map[string]*websocket.Conn)

// 处理ws请求
func WsHandler(w http.ResponseWriter, r *http.Request) {
	var conn *websocket.Conn
	var err error
	conn, err = wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("连接出错：", err)
		return
	} else {
		remoteAddr := conn.RemoteAddr().String()
		fmt.Println("连上了，地址：", remoteAddr)
		_, ok := connMap[remoteAddr]
		if !ok {
			connMap[remoteAddr] = conn
		}
		fmt.Println("当前连接总数：", len(connMap))
	}

}

func main() {
	go runMonitorProcessTicker()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.LoadHTMLGlob("../Views/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/monitor", func(c *gin.Context) {
		WsHandler(c.Writer, c.Request)
	})
	r.Run()
}

//监控进程数据
func runMonitorProcessTicker() {
	for range time.NewTicker(time.Second * 1).C {
		if len(connMap) > 0 {
			processInfo := Model.GetProcessInfo()
			//推送
			for k, conn := range connMap {
				err := conn.WriteJSON(processInfo)
				if err != nil {
					delete(connMap, k)
					fmt.Println("当前连接总数：", len(connMap))
					fmt.Println(conn.RemoteAddr().String(), "已断开")
				}
			}
		}
	}
}
