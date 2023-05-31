package app

import (
	"go-pixiv/model/illust"
	"sync"
	"time"
)

type Illusts struct {
	wait    time.Duration
	handles []HandleFunc
	Group   *Group
}

func (t *Illusts) handle(ids []string, fun func(index int, id string)) {
	wait := sync.WaitGroup{}
	for i := 0; i < len(ids); time.Sleep(t.wait) {
		wait.Add(1)
		index := i
		id := ids[i]
		go func() {
			defer wait.Done()
			fun(index, id)
		}()
		i++
	}
}

func (t *Illusts) Use(handles ...HandleFunc) {
	t.handles = append(t.handles, handles...)
}

func (t *Illusts) SetWaitTime(wait time.Duration) {
	t.wait = wait
}

func (t *Illusts) Detailed(ids []string, handles ...HandleFunc) (illusts []illust.Illust) {
	handles = append(t.handles, handles...)
	url := t.Group.engine.Url
	illusts = make([]illust.Illust, len(ids))
	t.handle(ids, func(index int, id string) {
		t.Group.Get(url.Illust(id), append([]HandleFunc{func(ctx *Context) {
			if ctx.Err == nil {
				ctx.Err = ctx.Json(&illusts[index])
			}
		}}, handles...)...)
	})
	return
}

func (t *Illusts) Pages(ids []string, handles ...HandleFunc) (illustsPages []illust.IllustPages) {
	handles = append(t.handles, handles...)
	url := t.Group.engine.Url
	illustsPages = make([]illust.IllustPages, len(ids))
	t.handle(ids, func(index int, id string) {
		t.Group.Get(url.IllustPages(id), append([]HandleFunc{func(ctx *Context) {
			if ctx.Err == nil {
				ctx.Err = ctx.Json(&illustsPages[index])
			}
		}}, handles...)...)
	})
	return
}
