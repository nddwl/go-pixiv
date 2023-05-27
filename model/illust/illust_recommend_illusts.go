package illust

import "time"

//IllustRecommendIllusts https://www.pixiv.net/ajax/illust/recommend/illusts?illust_ids[]=<illust_id>&lang=<lang>&version=<version>
type IllustRecommendIllusts struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illusts []struct {
			ID             string   `json:"id"`
			Title          string   `json:"title"`
			IllustType     int      `json:"illustType"`
			XRestrict      int      `json:"xRestrict"`
			Restrict       int      `json:"restrict"`
			Sl             int      `json:"sl"`
			URL            string   `json:"url"`
			Description    string   `json:"description"`
			Tags           []string `json:"tags"`
			UserID         string   `json:"userId"`
			UserName       string   `json:"userName"`
			Width          int      `json:"width"`
			Height         int      `json:"height"`
			PageCount      int      `json:"pageCount"`
			IsBookmarkable bool     `json:"isBookmarkable"`
			BookmarkData   struct {
				ID      string `json:"id"`      //ID 该收藏记录的唯一标识符,可以用于后续的修改或删除操作
				Private bool   `json:"private"` //Private 是否为私人收藏,如果为 true,则表示这个收藏仅对用户自己可见,如果为 false,则表示这个收藏为公开的
			} `json:"bookmarkData"` //BookmarkData 收藏数据
			Alt                     string `json:"alt"`
			TitleCaptionTranslation struct {
				WorkTitle struct {
					En string //En 根据用户的语言或lang来返回结果
					Ja string //Ja 默认
				} `json:"workTitle"` //WorkTitle 标题翻译
				WorkCaption struct {
					En string //En 根据用户的语言或lang来返回结果
					Ja string //Ja 默认
				} `json:"workCaption"` //WorkCaption 描述翻译
			} `json:"titleCaptionTranslation"` //作品标题和作品描述的翻译信息
			CreateDate      time.Time `json:"createDate"`
			UpdateDate      time.Time `json:"updateDate"`
			IsUnlisted      bool      `json:"isUnlisted"`
			IsMasked        bool      `json:"isMasked"`
			AiType          int       `json:"aiType"`
			ProfileImageURL string    `json:"profileImageUrl"`
			Type            string    `json:"type"`
			IsAdContainer   bool      `json:"isAdContainer"`
		} `json:"illusts"`
	} `json:"body"`
}
