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
	id, err := createUserGroup(usergroup)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	fmt.Println(id)
}
