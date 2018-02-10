package storage

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/youtangai/eniwa03/api/config"
)

var (
	// DataBase is struct for db
	DataBase *gorm.DB
)

const (
	//DBMS is database management system
	DBMS = "mysql"
)

func init() {
	DataBase = connection()
	//DataBase.SingularTable(true)
}

func connection() *gorm.DB {
	user := config.DBUser()
	pass := config.DBPasswd()
	host := config.DBHost()
	port := config.DBPort()
	dbName := config.DBName()
	connectionString := user + ":" + pass + "@" + "tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	db, err := gorm.Open(DBMS, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
