package model

type Dispatch struct {
	*Model
	TopicId int    `json:"topic_id"`
	Target  string `json:"target"`
}
