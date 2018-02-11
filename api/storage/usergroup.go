package storage

import (
	"log"
	"strconv"

	"github.com/youtangai/eniwa03/api/model"
)

const (
	JOIN   = 1
	REJECT = -1
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
	joinFlag := strconv.Itoa(usergroup.JoinFlag)

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
		select id, group_name, state, join_flag from user_groups 
		inner join groups on user_groups.group_id = groups.id
		where user_id = '` + userID + `'
	`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var joinFlag string
		err := rows.Scan(&(group.ID), &(group.GroupName), &(group.Status), &(joinFlag))
		if err != nil {
			return nil, err
		}
		if joinFlag != "-1" {
			if joinFlag == "0" {
				group.Status = "invite"
			} else {
				if group.Status == "0" {
					group.Status = "doing"
				} else if group.Status == "1" {
					group.Status = "success"
				} else {
					group.Status = "failed"
				}
			}
			groups = append(groups, group)
		}
	}
	return groups, nil
}

func CreateUserGroup(userid, groupid string) error {
	var usergroup model.UserGroup
	userID, err := strconv.Atoi(userid)
	if err != nil {
		return err
	}
	groupID, err := strconv.Atoi(groupid)
	if err != nil {
		return err
	}
	usergroup.UserID = userID
	usergroup.GroupID = groupID
	err = createUserGroup(usergroup)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserGroupJoin(userid, groupid, status string) error {
	var usergroup model.UserGroup
	userID, err := strconv.Atoi(userid)
	if err != nil {
		return err
	}
	groupID, err := strconv.Atoi(groupid)
	if err != nil {
		return err
	}
	usergroup.UserID = userID
	usergroup.GroupID = groupID

	joinStatus, err := strconv.Atoi(status)
	if err != nil {
		return err
	}
	usergroup.JoinFlag = joinStatus
	err = updateUserGroup(usergroup)
	if err != nil {
		return err
	}
	return nil
}

func GetIndividuals(groupid string) ([]model.Individual, error) {
	var individuals []model.Individual
	rows, err := DataBase.Query(`
		select id, name, current_price, goal_price from user_groups 
		inner join users on user_groups.user_id = users.id
		where group_id = '` + groupid + `'	
	`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var individual model.Individual
		err := rows.Scan(&(individual.UserID), &(individual.UserName), &(individual.Current), &(individual.Goal))
		if err != nil {
			return nil, err
		}
		individuals = append(individuals, individual)
	}
	return individuals, nil
}

func GetUserGroupByUseridGroupid(userid, groupid string) (model.UserGroup, string, error) {
	var usergroup model.UserGroup
	var userName string
	row := DataBase.QueryRow(`
		select user_id, group_id, goal_price, current_price, goal_desc, join_flag, name from user_groups
		inner join users on user_id = users.id
		where user_id = '` + userid + `' and group_id = '` + groupid + `'	
	`)
	err := row.Scan(&(usergroup.UserID), &(usergroup.GroupID), &(usergroup.GoalPrice), &(usergroup.CurrentPrice), &(usergroup.GoalDesc), &(usergroup.JoinFlag), &(userName))
	if err != nil {
		return usergroup, "", err
	}
	return usergroup, userName, nil
}

func SetGoalUserGroup(userid, groupid, price, desc string) error {
	result, err := DataBase.Exec(`
		update user_groups set goal_price = '` + price + `', goal_desc = '` + desc + `' where user_id = '` + userid + `' and group_id = '` + groupid + `'
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %v", result)
	return nil
}

func AddCurrentPrice(userid, groupid, price string) error {
	result, err := DataBase.Exec(`
		update user_groups set current_price = current_price + ` + price + ` where user_id = '` + userid + `' and group_id = '` + groupid + `'	
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %v", result)
	return nil
}
