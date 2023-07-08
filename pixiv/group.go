package pixiv

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Group struct {
	engine  *Engine
	handles []HandleFunc
	Url     Url
}

func (t *Group) do(req *http.Request, err error, handles []HandleFunc) *Context {
	ctx := &Context{
		Request:  req,
		Response: nil,
		Err:      err,
		engine:   t.engine,
		handles:  handles,
		index:    -1,
	}
	if req != nil && err == nil {
		req.Header.Set("Host", req.URL.Host)
		if req.Header.Get("User-Agent") == "" {
			req.Header.Set("User-Agent", t.engine.Config.UserAgent)
		}
		if req.Header.Get("Referer") == "" {
			req.Header.Set("Referer", "https://www.pixiv.net/")
		}
		resp, err1 := t.engine.client.Do(req)
		ctx.Err = err1
		ctx.Response = resp
		for _, v := range t.handles {
			v(ctx)
		}
	}
	return ctx
}

func (t *Group) Do(req *http.Request, err error, handles ...HandleFunc) {
	t.do(req, err, handles).Next()
}

func (t *Group) Get(url string, handles ...HandleFunc) {
	req, err := http.NewRequest("GET", url, nil)
	t.do(req, err, handles).Next()
}

func (t *Group) GetJson(url string, obj interface{}, handles ...HandleFunc) {
	req, err := http.NewRequest("GET", url, nil)
	ctx := t.do(req, err, handles)
	if ctx.Err == nil {
		ctx.Err = ctx.Json(&obj)
	}
	ctx.Next()
}

func (t *Group) Post(url string, contentType string, body io.Reader, handles ...HandleFunc) {
	req, err := http.NewRequest("POST", url, body)
	if err == nil {
		req.Header.Set("Content-Type", contentType)
	}
	t.do(req, err, handles).Next()
}

func (t *Group) PostJson(url string, data interface{}, handles ...HandleFunc) {
	b, err := json.Marshal(&data)
	var req *http.Request
	if err == nil {
		req, err = http.NewRequest("GET", url, bytes.NewReader(b))
	}
	t.do(req, err, handles).Next()
}
