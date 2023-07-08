package illust

//Pages https://www.pixiv.net/ajax/illust/<illust_id>/pages?lang=<lang>&version=<version>
type Pages struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    []struct {
		Urls struct {
			ThumbMini string `json:"thumb_mini"`
			Small     string `json:"small"`
			Regular   string `json:"regular"`
			Original  string `json:"original"`
		} `json:"urls"`
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"body"`
}
