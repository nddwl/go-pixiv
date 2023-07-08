package config

import (
	"net/http"
	"net/url"
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
