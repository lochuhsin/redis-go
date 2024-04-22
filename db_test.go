package rredis

import (
	"math/rand"
	"testing"
	"time"
)

func Test_BucketCreateDB(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	db := r.Intn(1000)

	bucket := newBucket(db)

	if err := bucket.create(db + 1); err == nil {
		t.Error("Error should be raised when creating database that is larger than initial count setting")
	}
	if err := bucket.create(db - 1); err != nil {
		t.Error(err)
	}
}

func Test_FlushDB(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	db := r.Intn(1000)

	bucket := newBucket(db)
	bucket.create(db - 1)
	if err := bucket.flush(db - 1); err != nil {
		t.Error(err)
	}
}

func Test_DBExist(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	db := r.Intn(1000)

	bucket := newBucket(db)
	bucket.create(db - 1)
	if !bucket.exist(db - 1) {
		t.Error("database should exist")
	}
}

func Test_DBDelete(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	db := r.Intn(1000)

	bucket := newBucket(db)
	bucket.create(db - 1)
	bucket.flush(db - 1)
	if bucket.exist(db - 1) {
		t.Error("database should not exist")
	}
}
