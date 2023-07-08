package tags

import "go-pixiv/model/utils"

//Tags https://www.pixiv.net/ajax/tags/illust/<illust_id>?lang=<lang>&version=<version>
type Tags struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Tags utils.Tags `json:"tags"`
	} `json:"body"`
}
