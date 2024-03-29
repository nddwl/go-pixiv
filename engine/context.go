package engine

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
)

type Context struct {
	Err      error
	Request  *http.Request
	Response *http.Response
	engine   *Engine
	handles  []HandleFunc
	index    int8
}

func (t *Context) Abort() {
	t.index = math.MaxInt8
}

func (t *Context) Next() {
	t.index++
	for t.index < int8(len(t.handles)) {
		t.handles[t.index](t)
		t.index++
	}
	if t.Response != nil && !t.Response.Close {
		t.Response.Body.Close()
	}
}

func (t *Context) Byte() (b []byte, err error) {
	b, err = io.ReadAll(t.Response.Body)
	return
}

func (t *Context) Json(data interface{}) error {
	b, err := io.ReadAll(t.Response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &data)
	return err
}
