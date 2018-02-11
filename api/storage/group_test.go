package storage

import (
	"fmt"
	"testing"
	"time"

	"github.com/youtangai/eniwa03/api/model"
)

var (
	group model.Group
)

func TestCreateGroup(t *testing.T) {
	group.GroupName = "eniwa03"
	group.Start = time.Now()
	group.Dead = time.Now()
	group.State = 1
	id, err := createGroup(group)
	if err != nil {
		t.Fatalf("err = %#v", err)
	}
	group.ID = id
}

func TestReadGroups(t *testing.T) {
	groups, err := readGroups()
	if err != nil {
		t.Fatalf("err = %#v", err)
	}
	fmt.Printf("groups = %#v\n", groups)
}

func TestUPdateGroup(t *testing.T) {
	group.GroupName = "change"
	group.Start = time.Now()
	group.Dead = time.Now()
	group.State = 2
	err := updateGroup(group)
	if err != nil {
		t.Fatalf("err = %#v", err)
	}
}

func TestDeleteGroup(t *testing.T) {
	err := deleteGroup(group)
	if err != nil {
		t.Fatalf("err = %#v", err)
	}
}
