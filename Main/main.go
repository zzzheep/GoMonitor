package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
