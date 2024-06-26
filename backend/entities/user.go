package entities

import (
	"context"
)

type UserLoginPayload struct {
	Email    string `json:"email" bson:"email" validate:"email,required"`
	RoomCode string `json:"roomCode" bson:"roomCode" validate:"alphanum,required"`
}

type UserLoginResponse struct {
	Profile *UserData `json:"profile"`
}

type UserData struct {
	Id        string `json:"id" bson:"_id" validate:"alphanum"`
	FirstName string `json:"firstName" bson:"firstName" validate:"alpha"`
	LastName  string `json:"lastName" bson:"lastName" validate:"alpha"`
	Email     string `json:"email" bson:"email" validate:"email"`
	Image     string `json:"image" bson:"image" validate:"url"`
	*Audit    `bson:",inline"`
}

func (u *UserData) TableName() string {
	return "users"
}

type UserUpdatePayload struct {
	FirstName string `json:"firstName" validate:"alpha"`
	LastName  string `json:"lastName" validate:"omitempty,alpha"`
	Image     string `json:"image" validate:"omitempty,url"`
}

type UserSearchPayload struct {
	FirstName string `json:"firstName" validate:"omitempty,alpha"`
	LastName  string `json:"lastName" validate:"omitempty,alpha"`
	Email     string `json:"email" validate:"omitempty,email"`
	Skip      int64  `json:"skip" validate:"number"`
	Limit     int64  `json:"limit" validate:"number"`
}

type UserSearchResponse struct {
	Items []*UserData `json:"items"`
	Total int         `json:"total"`
}

type UserService interface {
	Login(ctx context.Context, payload *UserLoginPayload) (*UserLoginResponse, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*UserData, error)
	GetByEmail(ctx context.Context, email string) (*UserData, error)
	Update(ctx context.Context, id string, payload *UserUpdatePayload) (*UserData, error)
	Gets(ctx context.Context, payload *UserSearchPayload) (*UserSearchResponse, error)
}

type UserRepository interface {
	Create(ctx context.Context, payload *UserData) (*UserData, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*UserData, error)
	GetByIDs(ctx context.Context, ids []string) ([]*UserData, error)
	GetByEmail(ctx context.Context, email string) (*UserData, error)
	Update(ctx context.Context, id string, payload *UserData) (*UserData, error)
	Gets(ctx context.Context, filter interface{}, skip, limit int64) ([]*UserData, int64, error)
}
