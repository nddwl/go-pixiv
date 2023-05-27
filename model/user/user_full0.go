package user

//UserFull0 https://www.pixiv.net/ajax/user/<user_id>?full=0&lang=<lang>&version=<version>
type UserFull0 struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		UserID     string `json:"userId"`
		Name       string `json:"name"`
		Image      string `json:"image"`
		ImageBig   string `json:"imageBig"`
		Premium    bool   `json:"premium"`
		IsFollowed bool   `json:"isFollowed"`
		IsMypixiv  bool   `json:"isMypixiv"`
		IsBlocking bool   `json:"isBlocking"`
		Background struct {
			Repeat    bool   `json:"repeat"`
			Color     string `json:"color"`
			URL       string `json:"url"`
			IsPrivate bool   `json:"isPrivate"`
		} `json:"background"`
		SketchLiveID  string `json:"sketchLiveId"`
		Partial       int    `json:"partial"`
		AcceptRequest bool   `json:"acceptRequest"`
		SketchLives   []struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			URL           string `json:"url"`
			ThumbnailURL  string `json:"thumbnailUrl"`
			AudienceCount int    `json:"audienceCount"`
			IsR18         bool   `json:"isR18"`
			StreamerIds   []int  `json:"streamerIds"`
		} `json:"sketchLives"` //SketchLives
	} `json:"body"`
}
