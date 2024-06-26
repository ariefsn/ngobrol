package helper

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ariefsn/ngobrol/constants"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ParsePayload(r *http.Request, target interface{}) error {
	return json.NewDecoder(r.Body).Decode(&target)
}

func Capitalize(text string) string {
	return cases.Title(language.English, cases.NoLower).String(text)
}

func GetAuthEmailFromContext(ctx context.Context) string {
	if val, ok := ctx.Value(constants.AuthEmailKey).(string); ok {
		return val
	}

	return ""
}

func GetAuthRoomCodeFromContext(ctx context.Context) string {
	if val, ok := ctx.Value(constants.AuthRoomCodeKey).(string); ok {
		return val
	}

	return ""
}
