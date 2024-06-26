package directives

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ariefsn/ngobrol/constants"
	"github.com/ariefsn/ngobrol/entities"
)

type DirectiveFunc func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
type DirectiveProtectedFunc func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)

func Protected(userRepo entities.UserRepository) DirectiveProtectedFunc {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		authEmail := ctx.Value(constants.AuthEmailKey)
		if _, ok := authEmail.(string); !ok {
			return nil, errors.New("invalid auth email")
		}

		user, _ := userRepo.GetByEmail(ctx, authEmail.(string))

		if user == nil {
			return nil, errors.New("invalid auth email")
		}

		return next(ctx)
	}

}
