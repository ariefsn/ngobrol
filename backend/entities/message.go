package entities

import (
	"context"
)

type MessageCreatePayload struct {
	Message string `json:"message" bson:"message"`
}

type MessageData struct {
	Id      string `json:"id" bson:"_id" validate:"alphanum"`
	RoomId  string `json:"roomId" bson:"roomId" validate:"alphanum"`
	FromId  string `json:"fromId" bson:"fromId" validate:"alphanum"`
	Message string `json:"message" bson:"message"`
	IsNew   bool   `json:"isNew" bson:"isNew"`
	*Audit  `bson:",inline"`
}

func (u *MessageData) TableName() string {
	return "messages"
}

type MessageSearchPayload struct {
	IsNew bool  `json:"isNew" validate:"bool"`
	Skip  int64 `json:"skip" validate:"number"`
	Limit int64 `json:"limit" validate:"number"`
}

type MessageSearchResponse struct {
	Items []*MessageData `json:"items"`
	Total int            `json:"total"`
}

type MessageService interface {
	Create(ctx context.Context, payload *MessageCreatePayload) (*MessageData, error)
	GetByID(ctx context.Context, id string) (*MessageData, error)
	Gets(ctx context.Context, payload *MessageSearchPayload) (*MessageSearchResponse, error)
}

type MessageRepository interface {
	Create(ctx context.Context, payload *MessageData) (*MessageData, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*MessageData, error)
	GetByCode(ctx context.Context, code string) (*MessageData, error)
	Gets(ctx context.Context, filter interface{}, skip int64, limit int64) ([]*MessageData, int64, error)
}
