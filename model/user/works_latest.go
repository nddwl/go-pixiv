package user

import (
	"go-pixiv/model/utils"
)

//WorksLatest https://www.pixiv.net/ajax/user/<user_id>/works/latest?lang=<lang>&version=<version>
type WorksLatest struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illusts utils.IllustId `json:"illusts"`
		Novels  utils.IllustId `json:"novels"`
	} `json:"body"`
}
