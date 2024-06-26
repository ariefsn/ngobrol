package service

import (
	"context"
	"time"

	"github.com/ariefsn/ngobrol/entities"
	"github.com/ariefsn/ngobrol/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type messageService struct {
	messageRepo entities.MessageRepository
}

// Gets implements entities.MessageService.
func (r *messageService) Gets(ctx context.Context, payload *entities.MessageSearchPayload) (*entities.MessageSearchResponse, error) {
	skip := int64(0)
	limit := int64(1000)
	if payload.Skip > 0 {
		skip = payload.Skip
	}
	if payload.Limit > 0 {
		limit = payload.Limit
	}
	filter := bson.M{"roomId": helper.GetAuthRoomCodeFromContext(ctx)}
	if payload.IsNew {
		filter["isNew"] = true
	}
	res, total, _ := r.messageRepo.Gets(ctx, filter, skip, limit)

	return &entities.MessageSearchResponse{
		Items: res,
		Total: int(total),
	}, nil
}

// Create implements entities.MessageService.
func (r *messageService) Create(ctx context.Context, payload *entities.MessageCreatePayload) (*entities.MessageData, error) {
	return r.messageRepo.Create(ctx, &entities.MessageData{
		Id:      primitive.NewObjectID().Hex(),
		FromId:  helper.GetAuthEmailFromContext(ctx),
		RoomId:  helper.GetAuthRoomCodeFromContext(ctx),
		Message: payload.Message,
		IsNew:   true,
		Audit: &entities.Audit{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
}

// Delete implements entities.MessageService.
func (r *messageService) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetByCode implements entities.MessageService.
func (r *messageService) GetByCode(ctx context.Context, code string) (*entities.MessageData, error) {
	return r.messageRepo.GetByCode(ctx, code)
}

// GetByID implements entities.MessageService.
func (r *messageService) GetByID(ctx context.Context, id string) (*entities.MessageData, error) {
	panic("unimplemented")
}

// NewMessageService will create new an messageService object representation of entities.MessageService interface
func NewMessageService(messageRepo entities.MessageRepository) entities.MessageService {
	return &messageService{
		messageRepo: messageRepo,
	}
}
