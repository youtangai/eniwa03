package storage

import (
	"testing"
)

func TestCrateTable(t *testing.T) {
	err := createTable()
	if err != nil {
		t.Fatal(err)
	}
}
