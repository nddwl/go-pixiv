package user

import (
	"go-pixiv/model/utils"
)

//Illusts https://www.pixiv.net/ajax/user/<user_id>/illusts?ids[]=<illust_id>&lang=<lang>&version=<version>
type Illusts struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illust utils.Illust `json:"illust"`
	}
}
