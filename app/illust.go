package app

import "go-pixiv/model/illust"

type Illust struct {
	handles          []func(ctx *IContext)
	id               string
	illust           *illust.Illust
	pages            *illust.IllustPages
	recommend        *illust.IllustRecommend
	recommendIllusts *illust.IllustRecommendIllusts
	Group            *Group
}

func (t *Illust) Use(handles ...func(ctx *IContext)) {
	t.handles = append(t.handles, handles...)
}

func (t *Illust) ID() string {
	return t.id
}

func iCtxToCtx(illust *Illust, handles ...func(ctx *IContext)) []HandleFunc {
	handles = append(illust.handles, handles...)
	handleFunks := make([]HandleFunc, len(handles))
	for k, v := range handles {
		iCtx := &IContext{Illust: illust}
		handle := v
		handleFunks[k] = func(ctx *Context) {
			iCtx.Context = ctx
			handle(iCtx)
		}
	}
	return handleFunks
}

func (t *Illust) Detailed(handles ...func(ctx *IContext)) *illust.Illust {
	if t.illust != nil {
		return t.illust
	}
	t.Group.Get(t.Group.engine.Url.Illust(t.id), "application/json", iCtxToCtx(t, append([]func(ctx *IContext){func(ctx *IContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.illust)
		}
	}}, handles...)...)...)
	return t.illust
}

func (t *Illust) Pages(handles ...func(ctx *IContext)) *illust.IllustPages {
	if t.pages != nil {
		return t.pages
	}
	t.Group.Get(t.Group.engine.Url.IllustPages(t.id), "application/json", iCtxToCtx(t, append([]func(ctx *IContext){func(ctx *IContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.pages)
		}
	}}, handles...)...)...)
	return t.pages
}

//Recommend limit限制返回数据中illusts切片最大长度
func (t *Illust) Recommend(limit uint, handles ...func(ctx *IContext)) *illust.IllustRecommend {
	if t.recommend != nil {
		return t.recommend
	}
	t.Group.Get(t.Group.engine.Url.IllustRecommend(t.id, limit), "application/json", iCtxToCtx(t, append([]func(ctx *IContext){func(ctx *IContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.recommend)
		}
	}}, handles...)...)...)
	return t.recommend
}

func (t *Illust) RecommendIllusts(illustIds []string, handles ...func(ctx *IContext)) *illust.IllustRecommendIllusts {
	if t.recommendIllusts != nil {
		return t.recommendIllusts
	}
	t.Group.Get(t.Group.engine.Url.IllustRecommendIllusts(illustIds), "application/json", iCtxToCtx(t, append([]func(ctx *IContext){func(ctx *IContext) {
		if ctx.Err == nil {
			ctx.Err = ctx.Json(&t.recommendIllusts)
		}
	}}, handles...)...)...)
	return t.recommendIllusts
}

type IContext struct {
	*Context
	*Illust
}
