package model

import "time"

type User struct {
	ID       int `gorm:"primary_key;AUTO_INCREMENT"`
	Name     string
	Password string
}

type Group struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	GroupName string
	Start     time.Time
	End       time.Time
	State     int
}

type UserGroup struct {
	UserID       User  `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	GroupID      Group `gorm:"ForeignKey:GroupID;AssociationForeignKey:ID"`
	GoalPrice    int
	CurrentPrice int
	GoalDesc     string
	JoinFlag     bool
}
