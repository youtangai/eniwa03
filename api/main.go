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
		v1.POST("/invite_group", controller.InviteController)
		v1.GET("/details/:group_id", controller.DetailController)
		v1.GET("/set", controller.SetController)
		v1.POST("/setting", controller.SettingController)
		v1.POST("/charge", controller.ChargeController)
		v1.GET("/users", controller.UsersController)
		v1.POST("/invite", controller.InvitationController)
		v1.POST("/login", controller.CreateUserController)
	}
	router.Run(":" + PORT)
}
