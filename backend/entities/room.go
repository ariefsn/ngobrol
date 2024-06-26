package entities

import (
	"context"
)

type RoomCreatePayload struct {
	RoomCode string `json:"roomCode" bson:"roomCode" validate:"alphanum,required"`
	UserId   string `json:"userId" bson:"userId" validate:"email,required"`
}

type RoomData struct {
	Id     string   `json:"id" bson:"_id" validate:"alphanum"`
	Code   string   `json:"code" bson:"code" validate:"alphanum"`
	Image  string   `json:"image" bson:"image"`
	Users  []string `json:"users" bson:"users"`
	*Audit `bson:",inline"`
}

type RoomDataDetails struct {
	Id     string      `json:"id" bson:"_id" validate:"alphanum"`
	Code   string      `json:"code" bson:"code" validate:"alphanum"`
	Image  string      `json:"image" bson:"image"`
	Users  []*UserData `json:"users" bson:"users"`
	*Audit `bson:",inline"`
}

func (u *RoomData) TableName() string {
	return "rooms"
}

func (u *RoomDataDetails) TableName() string {
	return "rooms"
}

type RoomService interface {
	Create(ctx context.Context, payload *RoomCreatePayload) (*RoomDataDetails, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*RoomDataDetails, error)
	GetByCode(ctx context.Context, code string) (*RoomDataDetails, error)
	Update(ctx context.Context, id string, payload *RoomData) (*RoomData, error)
}

type RoomRepository interface {
	Create(ctx context.Context, payload *RoomData) (*RoomDataDetails, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*RoomDataDetails, error)
	GetByCode(ctx context.Context, code string) (*RoomDataDetails, error)
	Update(ctx context.Context, id string, payload *RoomData) (*RoomData, error)
}
