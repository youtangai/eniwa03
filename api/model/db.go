package model

import (
	"time"
)

//User is
type User struct {
	ID       int
	Name     string
	Password string
}

//Group is
type Group struct {
	ID        int
	GroupName string
	Start     time.Time
	Dead      time.Time
	State     int
}

//UserGroup is
type UserGroup struct {
	UserID       int
	GroupID      int
	GoalPrice    int
	CurrentPrice int
	GoalDesc     string
	JoinFlag     bool
}
