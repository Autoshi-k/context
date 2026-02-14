package context

import "context"

// for future probably idk

type Context context.Context

func Background() Context {
	return context.Background()
}
