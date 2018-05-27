package main

import (
	"GoMonitor/Model"
	"fmt"
	"log"
	"net/http"
	"runtime"
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
	runtime.GOMAXPROCS(runtime.NumCPU())
	go runMonitorProcessTicker()
	go runMonitorCpuTicker()
	go runMonitorNetTicker()
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
	defer func() {
		if result := recover(); result != nil {
			log.Println(result)
		}
	}()
	for range time.NewTicker(time.Second * 1).C {
		if len(connMap) > 0 {
			processInfo := Model.GetProcessInfo()
			//推送
			for _, conn := range connMap {
				conn.WriteJSON(gin.H{
					"type": "process",
					"data": processInfo,
				})
			}
		}
	}
}

//监控网路数据
func runMonitorNetTicker() {
	defer func() {
		if result := recover(); result != nil {
			log.Println(result)
		}
	}()
	for range time.NewTicker(time.Second * 1).C {
		if len(connMap) > 0 {
			netInfo := Model.GetNetInfo()
			//推送
			for k, conn := range connMap {
				err := conn.WriteJSON(gin.H{
					"type": "net",
					"data": netInfo,
				})
				if err != nil {
					delete(connMap, k)
					fmt.Println("当前连接总数：", len(connMap))
					fmt.Println(conn.RemoteAddr().String(), "已断开")
				}
			}
		}
	}
}

//监控cpu数据
func runMonitorCpuTicker() {
	defer func() {
		if result := recover(); result != nil {
			log.Println(result)
		}
	}()
	for range time.NewTicker(time.Second * 2).C {
		if len(connMap) > 0 {
			cpuInfo := Model.GetCpuInfo()
			//推送
			for _, conn := range connMap {
				conn.WriteJSON(gin.H{
					"type": "cpu",
					"data": cpuInfo,
				})
			}
		}
	}
}
