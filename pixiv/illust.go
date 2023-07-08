package pixiv

import (
	"go-pixiv/model/illust"
)

type Illust struct {
	handles []func(ctx *IContext)
	id      string
	Group   *Group
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

func (t *Illust) Detailed(handles ...func(ctx *IContext)) (illust illust.Illust) {
	url := t.Group.Url.Illust(t.id)
	t.Group.GetJson(url, &illust, iCtxToCtx(t, handles...)...)
	return
}

func (t *Illust) Pages(handles ...func(ctx *IContext)) (pages illust.Pages) {
	url := t.Group.Url.IllustPages(t.id)
	t.Group.GetJson(url, &pages, iCtxToCtx(t, handles...)...)
	return
}

func (t *Illust) Recommend(limit int, handles ...func(ctx *IContext)) (recommend *illust.Recommend) {
	url := t.Group.Url.IllustRecommend(t.id, limit)
	t.Group.GetJson(url, recommend, iCtxToCtx(t, handles...)...)
	return
}

func (t *Illust) RecommendIllusts(illustIds []string, handles ...func(ctx *IContext)) (recommend *illust.RecommendIllusts) {
	url := t.Group.Url.IllustRecommendIllusts(illustIds)
	t.Group.GetJson(url, &recommend, iCtxToCtx(t, handles...)...)
	return
}

type IContext struct {
	*Context
	*Illust
}
