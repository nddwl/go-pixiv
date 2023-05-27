package config

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	Cookie    string `json:"cookie"`
	Proxy     string `json:"proxy"`
	UserAgent string `json:"user_agent"`
	ProxyUrl  *url.URL
}

func (t *Config) Cookies() (c []*http.Cookie) {
	raw := strings.Split(t.Cookie, "; ")
	for _, cookie := range raw {
		cookie = strings.TrimSpace(cookie)
		if i := strings.IndexByte(cookie, '='); i != -1 {
			c = append(c, &http.Cookie{Name: cookie[:i], Value: cookie[i+1:]})
		}
	}
	return
}

func (t *Config) Save() {
	f, err := os.Open("./config.json")
	if err != nil {
		if os.IsNotExist(err) {
			b, _ := json.MarshalIndent(t, "", "	")
			f, _ = os.Create("./config.json")
			f.Write(b)
			f.Close()
		}
		panic(err)
	} else {
		b, _ := json.MarshalIndent(t, "", "	")
		f.Write(b)
		f.Close()
	}
}
