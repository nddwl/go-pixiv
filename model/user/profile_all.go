package user

//ProfileAll https://www.pixiv.net/ajax/user/<user_id>/profile/all?lang=<lang>&version=<version>
type ProfileAll struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illusts       []string    `json:"illust"`
		Manga         []string    `json:"manga"`
		Novels        []string    `json:"novels"`
		MangaSeries   MangaSeries `json:"mangaSeries"`
		NovelSeries   NovelSeries `json:"novelSeries"`
		Pickup        Pickup      `json:"pickup"`
		BookmarkCount struct {
			Public struct {
				Illust int `json:"illust"` //Works 公开插画数量
				Novel  int `json:"novel"`  //Novel 公开小说数量
			} `json:"public"` //Public 公开作品数量
			Private struct {
				Illust int `json:"illust"` //Works 私人插画数量
				Novel  int `json:"novel"`  //Novel 私人小说数量
			} `json:"private"` //Private 私人作品数量
		} `json:"bookmarkCount"` // 作品数量
		ExternalSiteWorksStatus struct {
			Booth    bool `json:"booth"`    //Booth 是否有与Booth相关
			Sketch   bool `json:"sketch"`   //Sketch 是否有与Sketch相关
			VroidHub bool `json:"vroidHub"` //VroidHub 是否有与VroidHub相关
		} `json:"externalSiteWorksStatus"` //ExternalSiteWorksStatus 外部站点作品状态
		Request struct {
			ShowRequestTab     bool `json:"showRequestTab"`     //ShowRequestTab 是否显示请求选项卡
			ShowRequestSentTab bool `json:"showRequestSentTab"` //ShowRequestSentTab 是否显示已发送的请求选项卡
			PostWorks          struct {
				Artworks []string `json:"artworks"` //Artworks 已请求的艺术作品
				Novels   []string `json:"novels"`   //Novels 已请求的小说作品
			} `json:"postWorks"` //PostWorks 已请求的作品数据
		} `json:"request"` //Request 请求数据
	} `json:"body"`
}
