package storage

import (
	"fmt"
	"testing"

	"github.com/youtangai/eniwa03/api/model"
)

var (
	usergroup model.UserGroup
)

func TestCreateUserGroup(t *testing.T) {
	usergroup.UserID = 10
	usergroup.GroupID = 4
	err := createUserGroup(usergroup)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
}

func TestReadUserGroups(t *testing.T) {
	usergroups, err := readUserGroups()
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	fmt.Printf("%v", usergroups)
}

func TestUpdateUserGroup(t *testing.T) {
	usergroup.UserID = 10
	usergroup.GroupID = 4
	usergroup.GoalPrice = 1000
	usergroup.CurrentPrice = 200
	usergroup.GoalDesc = "ipad買う"
	usergroup.JoinFlag = true
	err := updateUserGroup(usergroup)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
}
