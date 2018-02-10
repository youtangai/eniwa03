package model

import (
	"time"
)

type User struct {
	ID       uint
	Name     string
	Password string
}

type Group struct {
	ID        uint
	GroupName string
	Start     time.Time
	End       time.Time
	State     int
}

type UserGroup struct {
	UserID       uint
	GroupID      uint
	GoalPrice    int
	CurrentPrice int
	GoalDesc     string
	JoinFlag     bool
}
