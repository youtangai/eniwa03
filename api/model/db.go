package model

import (
	"time"
)

type User struct {
	ID       int
	Name     string
	Password string
}

type Group struct {
	ID        int
	GroupName string
	Start     time.Time
	End       time.Time
	State     int
}

type UserGroup struct {
	UserID       int
	GroupID      int
	GoalPrice    int
	CurrentPrice int
	GoalDesc     string
	JoinFlag     bool
}
