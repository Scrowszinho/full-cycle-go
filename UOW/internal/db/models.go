// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql"
)

type Category struct {
	ID   int32
	Name string
}

type Course struct {
	ID         int32
	Name       string
	CategoryID sql.NullInt32
}
