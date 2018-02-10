package storage

import (
	"fmt"
	"testing"

	"github.com/youtangai/eniwa03/api/model"
)

var (
	user model.User
)

func TestCreateUser(t *testing.T) {
	user.Name = "yota"
	user.Password = "1995"
	id, err := createUser(user)
	if err != nil {
		t.Fatalf("err = %#v\n", err)
	}
	user.ID = id
}

func TestReadUsers(t *testing.T) {
	users, err := readUsers()
	if err != nil {
		t.Fatalf("err = %#v\n", err)
	}
	fmt.Printf("users = %#v\n", users)
}

func TestUpdateUser(t *testing.T) {
	user.Name = "miyoshi"
	user.Password = "0116"
	err := updateUser(user)
	if err != nil {
		t.Fatalf("err = %#v", err)
	}
	fmt.Printf("user = %#v", user)
}

func TestDeleteUser(t *testing.T) {
	err := deleteUser(user)
	if err != nil {
		t.Fatalf("err = %+v\n", err)
	}
}
