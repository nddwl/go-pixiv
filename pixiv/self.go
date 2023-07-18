package pixiv

import (
	"go-pixiv/engine"
	"go-pixiv/model/self"
)

type Self struct {
	handles []func(ctx *SContext)
	self    *self.Detailed
	Group   *engine.Group
}

func sCtxToCtx(self *Self, handles ...func(ctx *SContext)) []engine.HandleFunc {
	handles = append(self.handles, handles...)
	handleFunks := make([]engine.HandleFunc, len(handles))
	for k, v := range handles {
		sCtx := &SContext{Self: self}
		handle := v
		handleFunks[k] = func(ctx *engine.Context) {
			sCtx.Context = ctx
			handle(sCtx)
		}
	}
	return handleFunks
}

func (t *Self) Use(handles ...func(ctx *SContext)) {
	t.handles = append(t.handles, handles...)
}

func (t *Self) Info(handles ...func(ctx *SContext)) *self.Detailed {
	if t.self != nil {
		return t.self
	}
	url := t.Group.Url.Self()
	t.Group.Get(url, sCtxToCtx(t, append([]func(ctx *SContext){func(ctx *SContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.self)
		}
	}}, handles...)...)...)
	return t.self
}

type SContext struct {
	*Self
	*engine.Context
}
