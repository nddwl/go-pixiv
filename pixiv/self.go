package pixiv

import "go-pixiv/model/self"

type Self struct {
	handles []func(ctx *SContext)
	self    *self.Self
	Group   *Group
}

func sCtxToCtx(self *Self, handles ...func(ctx *SContext)) []HandleFunc {
	handles = append(self.handles, handles...)
	handleFunks := make([]HandleFunc, len(handles))
	for k, v := range handles {
		sCtx := &SContext{Self: self}
		handle := v
		handleFunks[k] = func(ctx *Context) {
			sCtx.Context = ctx
			handle(sCtx)
		}
	}
	return handleFunks
}

func (t *Self) Use(handles ...func(ctx *SContext)) {
	t.handles = append(t.handles, handles...)
}

func (t *Self) Info(handles ...func(ctx *SContext)) *self.Self {
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
	*Context
}
