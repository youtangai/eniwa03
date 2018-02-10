package storage

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	"github.com/youtangai/eniwa03/api/config"
	"github.com/youtangai/eniwa03/api/model"
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

func createTable() error {
	DataBase.AutoMigrate(&model.User{}, &model.Group{}, &model.UserGroup{})
	if !DataBase.HasTable(&model.User{}) {
		return errors.New("database:User table not created")
	}
	if !DataBase.HasTable(&model.Group{}) {
		return errors.New("database:Group table not created")
	}
	if !DataBase.HasTable(&model.UserGroup{}) {
		return errors.New("database:UserGroup table not created")
	}
	return nil
}
