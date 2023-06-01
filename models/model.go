package model

const (
	OK = iota
	DELETED
)

type Model struct {
	Id        int  `json:"id"`
	IsDeleted int8 `json:"is_deleted"`
}
