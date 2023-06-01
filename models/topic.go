package model

const (
	AUTO = iota
	MANUAL
)

type Topic struct {
	*Model
	Name string `json:"name"`
	Type int8   `json:"type"`
}
