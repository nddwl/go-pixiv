package pixiv

import "go-pixiv/model/user"

type User struct {
	uid            string
	simple         *user.Simple
	detailed       *user.Detailed
	illusts        *user.Illusts
	profileAll     *user.ProfileAll
	profileIllusts *user.ProfileIllusts
	profileTop     *user.ProfileTop
	worksLatest    *user.WorksLatest
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

func (t *User) Simple(handles ...func(ctx *UContext)) *user.Simple {
	if t.simple != nil {
		return t.simple
	}
	url := t.Group.Url.User(t.uid, 0)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.simple)
		}
	}}, handles...)...)...)
	return t.simple
}

func (t *User) Detailed(handles ...func(ctx *UContext)) *user.Detailed {
	if t.detailed != nil {
		return t.detailed
	}
	url := t.Group.Url.User(t.uid, 1)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.detailed)
		}
	}}, handles...)...)...)
	return t.detailed
}

func (t *User) ProfileAll(handles ...func(ctx *UContext)) *user.ProfileAll {
	if t.profileAll != nil {
		return t.profileAll
	}
	url := t.Group.Url.UserProfileAll(t.uid)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.profileAll)
		}
	}}, handles...)...)...)
	return t.profileAll
}

func (t *User) Illusts(ids []string, handles ...func(ctx *UContext)) *user.Illusts {
	if t.illusts != nil {
		return t.illusts
	}
	url := t.Group.Url.UserIllusts(t.uid, ids)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.illusts)
		}
	}}, handles...)...)...)
	return t.illusts
}

func (t *User) ProfileIllusts(ids []string, handles ...func(ctx *UContext)) *user.ProfileIllusts {
	if t.profileIllusts != nil {
		return t.profileIllusts
	}
	url := t.Group.Url.UserProfileIllusts(t.uid, ids)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.profileIllusts)
		}
	}}, handles...)...)...)
	return t.profileIllusts
}

func (t *User) ProfileTop(handles ...func(ctx *UContext)) *user.ProfileTop {
	if t.profileTop != nil {
		return t.profileTop
	}
	url := t.Group.Url.UserProfileTop(t.uid)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.profileTop)
		}
	}}, handles...)...)...)
	return t.profileTop
}

func (t *User) WorksLatest(handles ...func(ctx *UContext)) *user.WorksLatest {
	if t.worksLatest != nil {
		return t.worksLatest
	}
	url := t.Group.Url.UserWorksLatest(t.uid)
	t.Group.Get(url, uCtxToCtx(t, append([]func(ctx *UContext){func(ctx *UContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.worksLatest)
		}
	}}, handles...)...)...)
	return t.worksLatest
}
