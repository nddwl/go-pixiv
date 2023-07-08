package pixiv

import (
	"go-pixiv/model/illust"
	"sync"
	"time"
)

type Illusts struct {
	wait    time.Duration
	handles []func(ctx *IContext)
	Group   *Group
}

func illustGo(ids []string, wait time.Duration, fun func(index int, id string)) {
	wg := sync.WaitGroup{}
	for i := 0; i < len(ids); time.Sleep(wait) {
		wg.Add(1)
		index := i
		id := ids[i]
		go func() {
			defer wg.Done()
			fun(index, id)
		}()
		i++
	}
}

func (t *Illusts) Use(handles ...func(ctx *IContext)) {
	t.handles = append(t.handles, handles...)
}

func (t *Illusts) SetWaitTime(wait time.Duration) {
	t.wait = wait
}

func (t *Illusts) Detailed(ids []string, handles ...func(ctx *IContext)) (detailed []illust.Illust) {
	detailed = make([]illust.Illust, len(ids))
	illustGo(ids, t.wait, func(index int, id string) {
		url := t.Group.Url.Illust(id)
		t.Group.GetJson(url, &detailed[index], iCtxToCtx(&Illust{id: id, Group: t.Group}, handles...)...)
	})
	return
}

func (t *Illusts) Pages(ids []string, handles ...func(ctx *IContext)) (pages []illust.Pages) {
	pages = make([]illust.Pages, len(ids))
	illustGo(ids, t.wait, func(index int, id string) {
		url := t.Group.Url.IllustPages(id)
		t.Group.GetJson(url, &pages[index], iCtxToCtx(&Illust{id: id, Group: t.Group}, handles...)...)
	})
	return
}

func (t *Illusts) Recommend(ids []string, limit int, handles ...func(ctx *IContext)) (recommend []illust.Recommend) {
	recommend = make([]illust.Recommend, len(ids))
	illustGo(ids, t.wait, func(index int, id string) {
		url := t.Group.Url.IllustRecommend(id, limit)
		t.Group.GetJson(url, &recommend[index], iCtxToCtx(&Illust{id: id, Group: t.Group}, handles...)...)
	})
	return
}

func (t *Illusts) RecommendIllusts(ids []string, handles ...func(ctx *IContext)) (recommend illust.RecommendIllusts) {
	url := t.Group.Url.IllustRecommendIllusts(ids)
	t.Group.GetJson(url, &recommend, iCtxToCtx(&Illust{id: ids[0], Group: t.Group}, handles...)...)
	return
}
