package middlewares

import (
	"context"
	"net/http"

	"github.com/ariefsn/ngobrol/constants"
	"github.com/ariefsn/ngobrol/entities"
	"github.com/ariefsn/ngobrol/helper"
)

func Inject(env helper.Env, userService entities.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Inject Writer
			ctx := context.WithValue(r.Context(), constants.WriterCtxKey, w)
			r = r.WithContext(ctx)

			// Inject Email
			email := r.Header.Get(string(constants.HeaderXEmail))
			if email != "" {
				newCtx := context.WithValue(r.Context(), constants.AuthEmailKey, email)
				r = r.WithContext(newCtx)
			}

			// Inject Code
			code := r.Header.Get(string(constants.HeaderXRoomCode))
			if code != "" {
				newCtx := context.WithValue(r.Context(), constants.AuthRoomCodeKey, code)
				r = r.WithContext(newCtx)
			}

			next.ServeHTTP(w, r)
		})
	}
}
