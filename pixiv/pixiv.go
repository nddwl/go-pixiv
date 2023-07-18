package pixiv

import (
	"go-pixiv/engine"
	"time"
)

type Pixiv struct {
	engine *engine.Engine
}

func New() *Pixiv {
	return &Pixiv{
		engine: engine.New(),
	}
}

func (t *Pixiv) Illust(id string, handles ...func(ctx *IContext)) (illust *Illust) {
	illust = &Illust{
		id:    id,
		Group: t.engine.Group,
	}
	if handles != nil {
		illust.Group.Do(nil, nil, iCtxToCtx(illust, handles...)...)
	}
	return
}

func (t *Pixiv) Illusts(handles ...func(ctx *IContext)) Illusts {
	illusts := Illusts{
		wait:    time.Millisecond * 200,
		handles: handles,
		Group:   t.engine.Group,
	}
	return illusts
}

func (t *Pixiv) Self(handles ...func(ctx *SContext)) (self *Self) {
	self = &Self{
		Group: t.engine.Group,
	}
	if handles != nil {
		self.Group.Do(nil, nil, sCtxToCtx(self, handles...)...)
	}
	return
}

func (t *Pixiv) User(uid string, handles ...func(ctx *UContext)) (user *User) {
	user = &User{
		uid:   uid,
		Group: t.engine.Group,
	}
	if handles != nil {
		user.Group.Do(nil, nil, uCtxToCtx(user, handles...)...)
	}
	return
}

func (t *Pixiv) Use(handles ...engine.HandleFunc) {
	t.engine.Use(handles...)
}

func (t *Pixiv) Save() {
	t.engine.Save()
}
