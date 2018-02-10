package storage

import (
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/youtangai/eniwa03/api/config"
	"github.com/youtangai/eniwa03/api/model"
)

var (
	// DataBase is struct for db
	DataBase *sql.DB
)

const (
	//DBMS is database management system
	DBMS = "mysql"
)

func init() {
	DataBase = connection()
}

func connection() *sql.DB {
	user := config.DBUser()
	pass := config.DBPasswd()
	host := config.DBHost()
	port := config.DBPort()
	dbName := config.DBName()
	connectionString := user + ":" + pass + "@" + "tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	db, err := sql.Open(DBMS, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

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
