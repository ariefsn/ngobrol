package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"time"

	"github.com/ariefsn/ngobrol/constants"
	"github.com/ariefsn/ngobrol/entities"
)

// SendMessage is the resolver for the sendMessage field.
func (r *mutationResolver) SendMessage(ctx context.Context, input entities.MessageCreatePayload) (*entities.MessageData, error) {
	return r.MessageService.Create(ctx, &input)
}

// GetMessages is the resolver for the getMessages field.
func (r *queryResolver) GetMessages(ctx context.Context, input entities.MessageSearchPayload) (*entities.MessageSearchResponse, error) {
	res, err := r.MessageService.Gets(ctx, &input)
	return res, err
}

// SubNewMessage is the resolver for the subNewMessage field.
func (r *subscriptionResolver) SubNewMessage(ctx context.Context, code string) (<-chan *entities.MessageData, error) {
	ch := make(chan *entities.MessageData)

	go func() {
		defer close(ch)

		for {
			time.Sleep(1 * time.Second)

			res, _ := r.MessageService.Gets(context.WithValue(ctx, constants.AuthRoomCodeKey, code), &entities.MessageSearchPayload{
				IsNew: true,
				Skip:  0,
				Limit: 1,
			})

			var msg *entities.MessageData

			if res.Total > 0 {
				msg = res.Items[0]
				// r.MessageService.
			}

			select {
			case <-ctx.Done():
				return
			case ch <- msg:
			}
		}
	}()

	return ch, nil
}