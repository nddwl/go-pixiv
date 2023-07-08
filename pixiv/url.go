package pixiv

import (
	"strconv"
	"strings"
)

type Url struct{}

func (t Url) Self() string {
	return "https://www.pixiv.net/ajax/user/self"
}

func (t Url) User(uid string, full int) string {
	return "https://www.pixiv.net/ajax/user/" + uid + "?full=" + strconv.Itoa(full)
}

func (t Url) UserProfileAll(uid string) string {
	return "https://www.pixiv.net/ajax/user/" + uid + "/profile/all"
}

func (t Url) UserIllusts(uid string, ids []string) string {
	return "https://www.pixiv.net/ajax/user/" + uid + "/illusts?ids[]=" + strings.Join(ids, "&ids[]=")
}

func (t Url) UserProfileIllusts(uid string, ids []string) string {
	return "https://www.pixiv.net/ajax/user/" + uid + "/profile/illusts?ids[]=" + strings.Join(ids, "&ids[]=")
}

func (t Url) UserProfileTop(uid string) string {
	return "https://www.pixiv.net/ajax/user/" + uid + "/profile/top"
}

func (t Url) UserWorksLatest(uid string) string {
	return "https://www.pixiv.net/ajax/user/" + uid + "/works/latest"
}

func (t Url) Illust(id string) string {
	return "https://www.pixiv.net/ajax/illust/" + id
}

func (t Url) IllustPages(id string) string {
	return "https://www.pixiv.net/ajax/illust/" + id + "/pages"
}

func (t Url) IllustRecommend(id string, limit int) string {
	return "https://www.pixiv.net/ajax/illust/" + id + "/recommend/init?limit=" + strconv.Itoa(limit)
}

func (t Url) IllustRecommendIllusts(illustIds []string) string {
	return "https://www.pixiv.net/ajax/illust/recommend/illusts?illust_ids[]=" + strings.Join(illustIds, "&illust_ids[]=")
}
