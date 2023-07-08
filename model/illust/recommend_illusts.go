package illust

import (
	"go-pixiv/model/utils"
)

//RecommendIllusts https://www.pixiv.net/ajax/illust/recommend/illusts?illust_ids[]=<illust_id>&lang=<lang>&version=<version>
type RecommendIllusts struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illusts []utils.Work `json:"illusts"`
	} `json:"body"`
}
