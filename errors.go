package rredis

import "fmt"

type DatabaseError struct {
	detail string
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf(e.detail)
}

type DatabaseExistsError struct {
	dbName int
}

func (e DatabaseExistsError) Error() string {
	return fmt.Sprintf("Database %v Exists", e.dbName)
}

type DatabaseNotFoundError struct {
	dbName int
}

func (e DatabaseNotFoundError) Error() string {
	return fmt.Sprintf("Database %v Not Found", e.dbName)
}

type DatabaseOptionError struct {
	detail string
}

func (e DatabaseOptionError) Error() string {
	return fmt.Sprintf("Database Option Error: %v", e.detail)
}

type DatabaseKeyError struct {
	key    string
	detail string
}

func (e DatabaseKeyError) Error() string {
	return fmt.Sprintf("Database key Error, Key: %v, Detail %v", e.key, e.detail)
}
