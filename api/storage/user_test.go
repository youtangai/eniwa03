package storage

import (
	"fmt"
	"testing"

	"github.com/youtangai/eniwa03/api/model"
)

func TestCreateUser(t *testing.T) {
	user := model.User{Name: "yota", Password: "1995"}
	err := createUser(user)
	if err != nil {
		t.Fatalf("err = %v\n", err)
	}
}

func TestReadUsers(t *testing.T) {
	users, err := readUsers()
	if err != nil {
		t.Fatalf("err = %+v\n", err)
	}
	fmt.Printf("users = %+v\n", users)
}

func TestDeleteUser(t *testing.T) {
	user := model.User{ID: 1}
	err := deleteUser(user)
	if err != nil {
		t.Fatalf("err = %+v\n", err)
	}
}
