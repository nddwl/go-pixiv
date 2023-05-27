package app

import (
	"fmt"
	"strings"
)

type Url struct {
	*Group
}

func (t Url) Self() string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/user/self?lang=%s", t.engine.lang)
}

//User full可选值为0或1
func (t Url) User(uid string, full int) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/user/%s?full=%d&lang=%s", uid, full, t.engine.lang)
}

func (t Url) UserProfileAll(uid string) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/user/%s/profile/all?lang=%s", uid, t.engine.lang)
}

func (t Url) UserIllusts(uid string, ids []string) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/user/%s/illusts?ids[]=%s&lang=%s", uid, strings.Join(ids, "&ids=[]"), t.engine.lang)
}

func (t Url) UserProfileIllusts(uid string, ids []string) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/user/%s/profile/illusts?ids[]=%s&lang=%s", uid, strings.Join(ids, "&ids=[]"), t.engine.lang)
}

func (t Url) UserProfileTop(uid string) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/user/%s/profile/top?lang=%s", uid, t.engine.lang)
}

func (t Url) UserWorksLatest(uid string) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/user/%s/works/latest?lang=%s", uid, t.engine.lang)
}

func (t Url) Illust(id string) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/illust/%s?lang=%s", id, t.engine.lang)
}

func (t Url) IllustPages(id string) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/illust/%s/pages?lang=%s", id, t.engine.lang)
}

func (t Url) IllustRecommend(id string, limit uint) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/illust/%s/recommend/init?limit=%d&lang=%s", id, limit, t.engine.lang)
}

func (t Url) IllustRecommendIllusts(illustIds []string) string {
	return fmt.Sprintf("https://www.pixiv.net/ajax/illust/recommend/illusts?illust_ids[]=%slang=%s", strings.Join(illustIds, "&illust_ids[]="), t.engine.lang)
}
