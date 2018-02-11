package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/eniwa03/api/model"
	"github.com/youtangai/eniwa03/api/storage"
)

//LoginController is
func LoginController(c *gin.Context) {
	var resp model.Login
	username := c.Query("name")
	password := c.Query("pass")
	var user model.User
	user.Name = username
	user.Password = password
	id, err := storage.GetUserIDByNamePass(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err = %v", err)
		return
	}
	resp.Status = "ok"
	resp.UserID = strconv.Itoa(id)
	c.JSON(http.StatusOK, resp)
}
