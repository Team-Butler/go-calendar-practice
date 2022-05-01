package cli

import (
	"context"
)

type Context struct {
	ctx context.Context

	// PORT
	// LOG FILE
	// ...
}

func NewContext(ctx context.Context) *Context {
	return &Context{ctx: ctx}
}
