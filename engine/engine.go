package engine

import (
	"encoding/json"
	"fmt"
	"go-pixiv/config"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"sync"
	"time"
)

type Engine struct {
	client *http.Client
	Config *config.Config
	pixiv  *url.URL
	rw     sync.RWMutex
	Group  *Group
}

func New() (e *Engine) {
	pixiv, _ := url.Parse("https://www.pixiv.net/")
	e = &Engine{
		Config: config.New(),
		pixiv:  pixiv,
		rw:     sync.RWMutex{},
	}
	e.init()
	return
}

func (t *Engine) init() {
	transport := &http.Transport{
		DialContext:           (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
		Proxy:                 http.ProxyURL(t.Config.ProxyUrl),
		TLSHandshakeTimeout:   5 * time.Second,
		MaxIdleConns:          100,
		IdleConnTimeout:       120 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
		ForceAttemptHTTP2:     true,
	}

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(t.pixiv, t.Config.Cookies())
	t.client = &http.Client{
		Transport:     transport,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       30 * time.Second,
	}
	t.Group = &Group{engine: t}
}

func (t *Engine) Save() {
	t.rw.Lock()
	cookies := t.client.Jar.Cookies(t.pixiv)
	str := ""
	for _, v := range cookies {
		str += v.String()
	}
	t.Config.Cookie = str
	f, err := os.Open("./config.json")
	if err != nil {
		if os.IsNotExist(err) {
			b, _ := json.MarshalIndent(t.Config, "", "	")
			f, _ = os.Create("./config.json")
			_, err1 := f.Write(b)
			if err1 != nil {
				fmt.Println(err1)
			}
			f.Close()
		}
		panic(err)
	} else {
		b, _ := json.MarshalIndent(t.Config, "", "	")
		_, err1 := f.Write(b)
		if err1 != nil {
			fmt.Println(err1)
		}
		f.Close()
	}
	t.rw.Unlock()
}

func (t *Engine) Use(handles ...HandleFunc) {
	t.Group.handles = append(t.Group.handles, handles...)
}
