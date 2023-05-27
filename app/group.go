package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Group struct {
	engine *Engine
}

func (t *Group) Do(req *http.Request, handles ...HandleFunc) {
	ctx := &Context{
		Request:  req,
		Response: nil,
		Err:      nil,
		engine:   t.engine,
		handles:  handles,
		index:    -1,
	}
	if req != nil {
		resp, err := t.engine.client.Do(req)
		ctx.Err = err
		ctx.Response = resp
	}
	t.engine.handle(ctx)
}

func (t *Group) Get(url string, contentType string, handles ...HandleFunc) {
	req, err := http.NewRequest("GET", url, nil)
	ctx := &Context{
		Request:  req,
		Response: nil,
		Err:      err,
		engine:   t.engine,
		handles:  handles,
		index:    -1,
	}
	if err == nil {
		req.Header.Set("Content-Type", contentType)
		resp, err1 := t.engine.client.Do(req)
		ctx.Err = err1
		ctx.Response = resp
	}
	t.engine.handle(ctx)
}

func (t *Group) Post(url string, contentType string, body io.Reader, handles ...HandleFunc) {
	req, err := http.NewRequest("POST", url, body)
	ctx := &Context{
		Request:  req,
		Response: nil,
		Err:      err,
		engine:   t.engine,
		handles:  handles,
		index:    -1,
	}
	if err == nil {
		req.Header.Set("Content-Type", contentType)
		resp, err1 := t.engine.client.Do(req)
		ctx.Err = err1
		ctx.Response = resp
	}
	t.engine.handle(ctx)
}

func (t *Group) PostJson(url string, data interface{}, handles ...HandleFunc) {
	b, err := json.Marshal(&data)
	ctx := &Context{
		Request:  nil,
		Response: nil,
		Err:      err,
		engine:   t.engine,
		handles:  handles,
		index:    -1,
	}
	if err == nil {
		req, err1 := http.NewRequest("GET", url, bytes.NewReader(b))
		ctx.Err = err1
		ctx.Request = req
		if err1 == nil {
			req.Header.Set("Content-Type", "application/json")
			resp, err2 := t.engine.client.Do(req)
			ctx.Err = err2
			ctx.Response = resp
		}
	}
	t.engine.handle(ctx)
}
