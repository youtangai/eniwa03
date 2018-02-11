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

const (
	TIME_FORMAT = "2016-01-02"
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

func InviteController(c *gin.Context) {
	userID := c.Query("user_id")
	groupID := c.Query("group_id")
	status := c.Query("status")
	err := storage.UpdateUserGroupJoin(userID, groupID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err = %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func DetailController(c *gin.Context) {
	groupid := c.Param("group_id")
	var detail model.GroupDetail
	//グループを１つ取得
	group, err := storage.GetGroupByID(groupid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err = %v", err)
		return
	}
	detail.State = group.State
	detail.Start = group.Start.Format(TIME_FORMAT)
	detail.Dead = group.Dead.Format(TIME_FORMAT)
	//ユーザと情報を取得
	individuals, err := storage.GetIndividuals(groupid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err = %v", err)
		return
	}
	detail.Individuals = individuals
	c.JSON(http.StatusOK, detail)
}

func SetController(c *gin.Context) {
	var resp model.GoalDetail
	userid := c.Query("user_id")
	groupid := c.Query("g_id")
	usergroup, err := storage.GetUserGroupByUseridGroupid(userid, groupid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err = %v", err)
		return
	}
	resp.CurrentPrice = usergroup.CurrentPrice
	resp.GoalPrice = usergroup.GoalPrice
	resp.Desc = usergroup.GoalDesc
	c.JSON(http.StatusOK, resp)
}

func SettingController(c *gin.Context) {
	userid := c.Query("user_id")
	groupid := c.Query("group_id")
	price := c.Query("price")
	desc := c.Query("description")
	err := storage.SetGoalUserGroup(userid, groupid, price, desc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("err = %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func ChargeController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
