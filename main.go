package context

import (
	"context"

	"github.com/google/uuid"
)

// for future probably idk

type Context context.Context

func Background() Context {
	ctx := context.WithValue(context.Background(), "uuid", uuid.NewString())
	return ctx
}
