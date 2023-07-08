package pixiv

import (
	"encoding/json"
	"fmt"
	"go-pixiv/config"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"sync"
	"time"
)

type Engine struct {
	client  *http.Client
	Config  *config.Config
	pixiv   *url.URL
	rw      sync.RWMutex
	Illusts *Illusts
	Group   *Group
}

func New() (e *Engine) {
	pixiv, _ := url.Parse("https://www.pixiv.net/")
	e = &Engine{
		Config: load(),
		pixiv:  pixiv,
		rw:     sync.RWMutex{},
	}
	e.init()
	e.initGroup()
	return
}

func NewWithConfig(cookie, proxy, ua string) (e *Engine) {
	pixiv, _ := url.Parse("https://www.pixiv.net/")
	e = &Engine{
		Config: &config.Config{
			Cookie:    cookie,
			Proxy:     proxy,
			UserAgent: ua,
		},
		pixiv: pixiv,
		rw:    sync.RWMutex{},
	}
	e.init()
	e.initGroup()
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
}

func (t *Engine) initGroup() {
	t.Group = &Group{engine: t}
	t.Illusts = &Illusts{
		wait:  time.Millisecond * 500,
		Group: t.Group,
	}
}

func load() (c *config.Config) {
	fmt.Println("读取配置...")
	f, err := os.Open("./config.json")
	if err != nil {
		if os.IsNotExist(err) {
			c = &config.Config{
				Cookie:    "",
				Proxy:     "",
				UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.39",
			}
			b, _ := json.MarshalIndent(c, "", "	")
			f, _ = os.Create("./config.json")
			f.Write(b)
			f.Close()
			log.Fatal("已生成配置文件,请填写...")
		}
		panic(err)
	}
	b, _ := io.ReadAll(f)
	f.Close()
	err = json.Unmarshal(b, &c)
	if err != nil {
		fmt.Println(err)
		log.Fatal("json配置文件错误...")
	}
	c.ProxyUrl, err = url.Parse(c.Proxy)
	if err != nil {
		log.Fatal("代理配置错误...")
	}
	fmt.Println("读取配置完成...")
	return
}

//Save 创建config.json保存cookie和proxy等数据
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
			f.Write(b)
			f.Close()
		}
		panic(err)
	} else {
		b, _ := json.MarshalIndent(t.Config, "", "	")
		f.Write(b)
		f.Close()
	}
	t.rw.Unlock()
}

func (t *Engine) Use(handles ...HandleFunc) {
	t.Group.handles = append(t.Group.handles, handles...)
}

func (t *Engine) User(uid string, handles ...func(ctx *UContext)) (user *User) {
	user = &User{
		uid:   uid,
		Group: t.Group,
	}
	if handles != nil {
		user.Group.Do(nil, nil, uCtxToCtx(user, handles...)...)
	}
	return
}

func (t *Engine) Self(handles ...func(ctx *SContext)) (self *Self) {
	self = &Self{
		Group: t.Group,
	}
	if handles != nil {
		self.Group.Do(nil, nil, sCtxToCtx(self, handles...)...)
	}
	return
}

func (t *Engine) Illust(id string, handles ...func(ctx *IContext)) (illust *Illust) {
	illust = &Illust{
		id:    id,
		Group: t.Group,
	}
	if handles != nil {
		illust.Group.Do(nil, nil, iCtxToCtx(illust, handles...)...)
	}
	return
}
