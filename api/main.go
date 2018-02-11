package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/eniwa03/api/controller"
)

const (
	//PORT is ポート番号
	PORT = "62070"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/login", controller.LoginController)
		v1.GET("/lists/:user_id", controller.ListsController)
		v1.POST("/make_group", controller.MakeGroupController)
	}
	router.Run(":" + PORT)
}
