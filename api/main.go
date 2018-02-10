package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	//PORT is ポート番号
	PORT = "62070"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":" + PORT)
}
