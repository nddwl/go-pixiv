package illust

import (
	"go-pixiv/model/utils"
)

//Recommend https://www.pixiv.net/ajax/illust/<illust_id>/recommend/init?limit=<num>&lang=<lang>&version=<version>
type Recommend struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illusts []utils.Work `json:"illust"`
		NextIds []string     `json:"nextIds"` //NextIds 下一批推荐作品ID列表
		Details Details      `json:"details"` //Details 匹配值
	} `json:"body"`
}
