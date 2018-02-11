package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

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

//ListsController is
func ListsController(c *gin.Context) {
	userID := c.Param("user_id")
	groups, err := storage.GetGroupsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err = %v", err)
		return
	}
	c.JSON(http.StatusOK, groups)
}

func MakeGroupController(c *gin.Context) {
	name := c.Query("name")
	date := c.Query("date")
	users := c.Query("users")
	users = strings.Trim(users, "[")
	users = strings.Trim(users, "]")
	userslice := strings.Split(users, ",")
	//グループ作成
	id, err := storage.CreateGroup(name, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err = %v", err)
		return
	}
	//招待作成
	idString := strconv.Itoa(id)
	for _, value := range userslice {
		err = storage.CreateUserGroup(idString, value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Printf("err = %v", err)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
