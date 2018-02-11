package storage

import (
	"log"
	"strconv"

	"github.com/youtangai/eniwa03/api/model"
)

func createGroup(group model.Group) (int, error) {
	start := group.Start.Format("2006-01-02 15:04:05")
	dead := group.Dead.Format("2006-01-02 15:04:05")
	state := strconv.Itoa(group.State)
	result, err := DataBase.Exec(`
		insert into groups(group_name, start, dead, state) values('` + group.GroupName + `', '` + start + `', '` + dead + `', '` + state + `')	
	`)
	if err != nil {
		return -1, err
	}
	log.Printf("result = %#v", result)
	id, err := result.LastInsertId()
	return int(id), nil
}

func readGroups() ([]model.Group, error) {
	var groups []model.Group
	rows, err := DataBase.Query(`
		select * from groups	
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var group model.Group
	for rows.Next() {
		err := rows.Scan(&(group.ID), &(group.GroupName), &(group.Start), &(group.Dead), &(group.State))
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func updateGroup(group model.Group) error {
	groupID := strconv.Itoa(group.ID)
	start := group.Start.Format("2006-01-02 15:04:05")
	dead := group.Dead.Format("2006-01-02 15:04:05")
	state := strconv.Itoa(group.State)
	result, err := DataBase.Exec(`
		update groups set group_name = '` + group.GroupName + `', start = '` + start + `', dead = '` + dead + `', state = ` + state + ` where id = '` + groupID + `'
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %#v", result)
	return nil
}

func deleteGroup(group model.Group) error {
	groupID := strconv.Itoa(group.ID)
	result, err := DataBase.Exec(`
		delete from groups where id = '` + groupID + `'	
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %#v", result)
	return nil
}
