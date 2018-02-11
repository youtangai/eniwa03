package storage

import (
	"log"
	"strconv"

	"github.com/youtangai/eniwa03/api/model"
)

func createUser(user model.User) (int, error) {
	result, err := DataBase.Exec(`
		insert into users(name, password) values('` + user.Name + `', '` + user.Password + `')
	`)
	if err != nil {
		return -1, err
	}
	log.Printf("result = %#v", result)
	id, err := result.LastInsertId()
	return int(id), nil
}

func readUsers() ([]model.User, error) {
	var users []model.User
	rows, err := DataBase.Query(`
		select * from users	
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {
		err := rows.Scan(&(user.ID), &(user.Name), &(user.Password))
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func updateUser(user model.User) error {
	userID := strconv.Itoa(user.ID)
	result, err := DataBase.Exec(`
		update users set name = '` + user.Name + `',password = '` + user.Password + `' where id = '` + userID + `'
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %#v", result)
	return nil
}

//id だけあればよい
func deleteUser(user model.User) error {
	userID := strconv.Itoa(user.ID)
	result, err := DataBase.Exec(`
		delete from users where id = '` + userID + `'
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %#v", result)
	return nil
}

func GetUserIDByNamePass(user model.User) (int, error) {
	var id int
	row := DataBase.QueryRow(`
		select id from users where name = '` + user.Name + `' and password = '` + user.Password + `'	
	`)
	err := row.Scan(&(id))
	if err != nil {
		return -1, err
	}
	return id, nil
}
