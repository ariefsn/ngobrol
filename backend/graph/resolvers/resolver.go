package resolvers

import "github.com/ariefsn/ngobrol/entities"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService    entities.UserService
	RoomService    entities.RoomService
	MessageService entities.MessageService
}
