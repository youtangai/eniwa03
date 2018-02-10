package storage

import (
	"testing"

	"github.com/youtangai/eniwa03/api/model"
)

func TestConnection(t *testing.T) {
	connection()
}

func TestCreateUser(t *testing.T) {
	user := model.User{Name: "yota", Password: "1995"}
	err := createUser(user)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
}
