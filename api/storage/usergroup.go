package storage

import (
	"log"
	"strconv"

	"github.com/youtangai/eniwa03/api/model"
)

func createUserGroup(usergroup model.UserGroup) error {
	userID := strconv.Itoa(usergroup.UserID)
	groupID := strconv.Itoa(usergroup.GroupID)
	result, err := DataBase.Exec(`
		insert into user_groups(user_id, group_id) values('` + userID + `', '` + groupID + `')	
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %#v", result)
	return nil
}

func readUserGroups() ([]model.UserGroup, error) {
	var usergroups []model.UserGroup
	rows, err := DataBase.Query(`
		select * from user_groups	
	`)
	if err != nil {
		return nil, err
	}
	var usergroup model.UserGroup
	for rows.Next() {
		err := rows.Scan(&(usergroup.UserID), &(usergroup.GroupID), &(usergroup.GoalPrice), &(usergroup.CurrentPrice), &(usergroup.GoalDesc), &(usergroup.JoinFlag))
		if err != nil {
			return nil, err
		}
		usergroups = append(usergroups, usergroup)
	}
	return usergroups, nil
}

func updateUserGroup(usergroup model.UserGroup) error {
	userID := strconv.Itoa(usergroup.UserID)
	groupID := strconv.Itoa(usergroup.GroupID)
	goalPrice := strconv.Itoa(usergroup.GoalPrice)
	currentPrice := strconv.Itoa(usergroup.CurrentPrice)
	joinFlag := "0"
	if usergroup.JoinFlag {
		joinFlag = "1"
	}

	result, err := DataBase.Exec(`
		update user_groups set goal_price = '` + goalPrice + `', current_price = '` + currentPrice + `', goal_desc = '` + usergroup.GoalDesc + `', join_flag = '` + joinFlag + `' where user_id = '` + userID + `' and group_id = '` + groupID + `'
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %v", result)
	return nil
}

func deleteUserGroup(usergroup model.UserGroup) error {
	userID := strconv.Itoa(usergroup.UserID)
	groupID := strconv.Itoa(usergroup.GroupID)
	result, err := DataBase.Exec(`
		delete from user_groups where user_id = '` + userID + `' and group_id = '` + groupID + `'
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %v", result)
	return nil
}

//GetGroupsByUserID is
func GetGroupsByUserID(userID string) ([]model.JoinedGroup, error) {
	var groups []model.JoinedGroup
	var group model.JoinedGroup
	rows, err := DataBase.Query(`
		select id, group_name, state from user_groups 
		inner join groups on user_groups.group_id = groups.id
		where user_id = '` + userID + `'
	`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&(group.ID), &(group.GroupName), &(group.Status))
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
