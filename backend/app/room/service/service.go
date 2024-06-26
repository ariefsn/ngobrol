package service

import (
	"context"
	"time"

	"github.com/ariefsn/ngobrol/entities"
)

type roomService struct {
	roomRepo entities.RoomRepository
}

// Update implements entities.RoomService.
func (r *roomService) Update(ctx context.Context, id string, payload *entities.RoomData) (*entities.RoomData, error) {
	return r.roomRepo.Update(ctx, id, payload)
}

// Create implements entities.RoomService.
func (r *roomService) Create(ctx context.Context, payload *entities.RoomCreatePayload) (*entities.RoomDataDetails, error) {
	existing, _ := r.roomRepo.GetByCode(ctx, payload.RoomCode)
	if existing == nil {
		existing, _ = r.roomRepo.Create(ctx, &entities.RoomData{
			Id:    payload.RoomCode,
			Code:  payload.RoomCode,
			Users: []string{payload.UserId},
			Audit: &entities.Audit{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
	}

	return existing, nil
}

// Delete implements entities.RoomService.
func (r *roomService) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetByCode implements entities.RoomService.
func (r *roomService) GetByCode(ctx context.Context, code string) (*entities.RoomDataDetails, error) {
	return r.roomRepo.GetByCode(ctx, code)
}

// GetByID implements entities.RoomService.
func (r *roomService) GetByID(ctx context.Context, id string) (*entities.RoomDataDetails, error) {
	panic("unimplemented")
}

// NewRoomService will create new an roomService object representation of entities.RoomService interface
func NewRoomService(roomRepo entities.RoomRepository) entities.RoomService {
	return &roomService{
		roomRepo: roomRepo,
	}
}
