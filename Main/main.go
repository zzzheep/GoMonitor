package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/monitor", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "monitorservice",
		})
	})
	r.GET("/monitor/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run()
}
