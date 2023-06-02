package model

// `create type consume_type as enum ('push', 'pull');`
type ConsumeType string

const (
	// PushConsumer, consume per message
	Push ConsumeType = "push"
	// PullConsumer, consume per user request (batch)
	Pull ConsumeType = "pull"
)

type Topic struct {
	*Model
	Name        string      `json:"name"`
	ConsumeType ConsumeType `json:"consume_type"`
}
