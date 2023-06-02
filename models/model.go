package model

import "database/sql"

const (
	OK = iota
	DELETED
)

type Model struct {
	Id        int          `json:"id"`
	IsDeleted sql.NullBool `json:"is_deleted"`
}
