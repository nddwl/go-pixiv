package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func New() (c *Config) {
	fmt.Println("读取配置...")
	f, err := os.Open("./config.json")
	if err != nil {
		if os.IsNotExist(err) {
			c = &Config{
				Cookie:    "",
				Proxy:     "",
				UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.39",
			}
			b, _ := json.MarshalIndent(c, "", "	")
			f, _ = os.Create("./config.json")
			_, err1 := f.Write(b)
			if err1 != nil {
				fmt.Println(err1)
			}
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
