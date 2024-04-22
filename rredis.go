package rredis

import "fmt"

func Init() {
	initBucket()
	fmt.Println("init complete")
}

func CreateDB(db int) error {
	return b.create(db)
}

func FlushDB(db int) error {
	return b.flush(db)
}

func DBExist(db int) bool {
	return b.exist(db)
}

func Get(k string, dbOption ...int) (string, error) {
	database, err := getDB(dbOption...)
	if err != nil {
		return "", err
	}
	return database.get(k)
}

func Set(k, v string, dbOption ...int) error {
	database, err := getDB(dbOption...)
	if err != nil {
		return err
	}
	return database.add(entry{k, v})
}

func Del(k string, dbOption ...int) error {
	database, err := getDB(dbOption...)
	if err != nil {
		return err
	}
	return database.delete(k)
}

func getDB(dbOption ...int) (*database, error) {
	var db int
	if len(dbOption) == 0 {
		db = 0
	} else if len(dbOption) > 1 {
		return nil, DatabaseOptionError{}
	}
	database, err := b.get(db)
	if err != nil {
		return nil, err
	}
	return database, nil
}
