package storage

import (
	"log"
	"strconv"

	"github.com/youtangai/eniwa03/api/model"
)

func createUser(user model.User) error {
	result, err := DataBase.Exec(`
		insert into users(name, password) values('` + user.Name + `', '` + user.Password + `')
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %+v", result)
	return nil
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

//id だけあればよい
func deleteUser(user model.User) error {
	userID := strconv.Itoa(user.ID)
	result, err := DataBase.Exec(`
		delete from users where id = '` + userID + `'
	`)
	if err != nil {
		return err
	}
	log.Printf("result = %+v", result)
	return nil
}
