package rredis

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func setup() {
	Init()
}
func teardown() {}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func Test_CreateDB_DBExist_FlushDB(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	db := rand.Intn(1024) // shift a range
	CreateDB(db)
	if DBExist(db) != true {
		t.Error(fmt.Printf("default database %v should exists", db))
	}
	err := FlushDB(db)
	if err != nil {
		t.Error(err)
	}
}

func Test_DefaultDBExist(t *testing.T) {
	if DBExist(0) != true {
		t.Error("default database 0 should exists")
	}
}

func Test_Add(t *testing.T) {
	k, expectV := "a", "b"
	err := Set(k, expectV)
	if err != nil {
		t.Error(err)
	}
	v, err := Get(k)
	if err != nil {
		t.Error(err)
	}
	if v != expectV {
		t.Error(v, expectV)
	}
}

func Test_AddDifferentDb(t *testing.T) {
	dbCount := 5
	k, expectV := "a", "b"
	for i := 1; i <= dbCount; i++ {
		CreateDB(i)
		Set(k, expectV, i)
	}

	for i := 1; i <= dbCount; i++ {
		v, err := Get(k)
		if err != nil {
			t.Error(err, i)
		}
		if v != expectV {
			t.Error(v, expectV, i)
		}
	}

	for i := 1; i <= dbCount; i++ {
		FlushDB(i)
	}
}
