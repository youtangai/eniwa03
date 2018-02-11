package storage

import (
	"log"
	"strconv"
	"time"

	"github.com/youtangai/eniwa03/api/model"
)

const (
	TIME_FORMAT = "2006-01-02"
)

func createGroup(group model.Group) (int, error) {
	start := group.Start.Format("2006-01-02 15:04:05")
	dead := group.Dead.Format("2006-01-02 15:04:05")
	result, err := DataBase.Exec(`
		insert into groups(group_name, start, dead) values('` + group.GroupName + `', '` + start + `', '` + dead + `')	
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
	// query := `update groups set group_name = '` + group.GroupName + `',start = '` + start + `',dead = '` + dead + `', state = '` + state + `' where id = '` + groupID + `'`
	// fmt.Println(query)
	// result, err := DataBase.Exec(query)
	result, err := DataBase.Exec(`
		update groups set group_name = '` + group.GroupName + `',start = '` + start + `',dead = '` + dead + `', state = '` + state + `' where id = '` + groupID + `'
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

func CreateGroup(name, date string) (int, error) {
	var group model.Group
	group.GroupName = name
	group.Start = time.Now()
	t, err := time.Parse(TIME_FORMAT, date)
	if err != nil {
		return -1, err
	}
	group.Dead = t
	id, err := createGroup(group)
	if err != nil {
		return -2, err
	}
	return id, nil
}
