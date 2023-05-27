package app

import "go-pixiv/model/user"

type User struct {
	uid            string
	full0          *user.UserFull0
	full1          *user.UserFull1
	illusts        *user.UserIllusts
	profileAll     *user.UserProfileAll
	profileIllusts *user.UserProfileIllusts
	profileTop     *user.UserProfileTop
	worksLatest    *user.UserWorksLatest
	handles        []func(ctx *UContext)
	Group          *Group
}

func uCtxToCtx(user *User, handles ...func(ctx *UContext)) []HandleFunc {
	handles = append(user.handles, handles...)
	handleFunks := make([]HandleFunc, len(handles))
	for k, v := range handles {
		uCtx := &UContext{User: user}
		handle := v
		handleFunks[k] = func(ctx *Context) {
			uCtx.Context = ctx
			handle(uCtx)
		}
	}
	return handleFunks
}

type UContext struct {
	*Context
	*User
}

func (t *User) Use(handles ...func(ctx *UContext)) {
	t.handles = append(t.handles, handles...)
}

func (t *User) UID() string {
	return t.uid
}

func (t *User) SimpleData(handles ...func(ctx *UContext)) *user.UserFull0 {
	if t.full0 != nil {
		return t.full0
	}
	url := t.Group.engine.Url.User(t.uid, 0)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.full0)
		}
	}}, handles...)...)...)
	return t.full0
}

func (t *User) DetailedData(handles ...func(ctx *UContext)) *user.UserFull1 {
	if t.full1 != nil {
		return t.full1
	}
	url := t.Group.engine.Url.User(t.uid, 1)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.full1)
		}
	}}, handles...)...)...)
	return t.full1
}

func (t *User) ProfileAll(handles ...func(ctx *UContext)) *user.UserProfileAll {
	if t.profileAll != nil {
		return t.profileAll
	}
	url := t.Group.engine.Url.UserProfileAll(t.uid)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.profileAll)
		}
	}}, handles...)...)...)
	return t.profileAll
}

func (t *User) Illusts(ids []string, handles ...func(ctx *UContext)) *user.UserIllusts {
	if t.illusts != nil {
		return t.illusts
	}
	url := t.Group.engine.Url.UserIllusts(t.uid, ids)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.illusts)
		}
	}}, handles...)...)...)
	return t.illusts
}

func (t *User) ProfileIllusts(ids []string, handles ...func(ctx *UContext)) *user.UserProfileIllusts {
	if t.profileIllusts != nil {
		return t.profileIllusts
	}
	url := t.Group.engine.Url.UserProfileIllusts(t.uid, ids)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.profileIllusts)
		}
	}}, handles...)...)...)
	return t.profileIllusts
}

func (t *User) ProfileTop(handles ...func(ctx *UContext)) *user.UserProfileTop {
	if t.profileTop != nil {
		return t.profileTop
	}
	url := t.Group.engine.Url.UserProfileTop(t.uid)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.profileTop)
		}
	}}, handles...)...)...)
	return t.profileTop
}

func (t *User) WorksLatest(handles ...func(ctx *UContext)) *user.UserWorksLatest {
	if t.worksLatest != nil {
		return t.worksLatest
	}
	url := t.Group.engine.Url.UserWorksLatest(t.uid)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.worksLatest)
		}
	}}, handles...)...)...)
	return t.worksLatest
}
