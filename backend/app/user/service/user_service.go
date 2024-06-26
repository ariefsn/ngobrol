package service

import (
	"context"
	"errors"
	"time"

	"github.com/ariefsn/ngobrol/entities"
)

type userService struct {
	userRepo entities.UserRepository
}

// Delete implements entities.UserService.
func (u *userService) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetByEmail implements entities.UserService.
func (u *userService) GetByEmail(ctx context.Context, email string) (*entities.UserData, error) {
	return u.userRepo.GetByEmail(ctx, email)
}

// GetByID implements entities.UserService.
func (u *userService) GetByID(ctx context.Context, id string) (*entities.UserData, error) {
	return u.userRepo.GetByID(ctx, id)
}

// Gets implements entities.UserService.
func (u *userService) Gets(ctx context.Context, payload *entities.UserSearchPayload) (*entities.UserSearchResponse, error) {
	res := &entities.UserSearchResponse{}

	users, total, err := u.userRepo.Gets(ctx, nil, payload.Skip, payload.Limit)
	if err != nil {
		return nil, err
	}
	res.Items = users
	res.Total = int(total)

	return res, nil
}

// Login implements entities.UserService.
func (u *userService) Login(ctx context.Context, payload *entities.UserLoginPayload) (*entities.UserLoginResponse, error) {
	user, _ := u.userRepo.GetByEmail(ctx, payload.Email)
	if user == nil {
		user, _ = u.userRepo.Create(ctx, &entities.UserData{
			Id:    payload.Email,
			Email: payload.Email,
			Audit: &entities.Audit{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
	}

	return &entities.UserLoginResponse{
		Profile: user,
	}, nil
}

// Update implements entities.UserService.
func (u *userService) Update(ctx context.Context, id string, payload *entities.UserUpdatePayload) (*entities.UserData, error) {
	existingUser, _ := u.GetByEmail(ctx, id)
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	existingUser.FirstName = payload.FirstName
	existingUser.LastName = payload.LastName
	existingUser.Image = payload.Image

	return u.userRepo.Update(ctx, id, existingUser)
}

// NewUserService will create new an userService object representation of entities.UserService interface
func NewUserService(userRepo entities.UserRepository) entities.UserService {
	return &userService{
		userRepo: userRepo,
	}
}
