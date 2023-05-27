package user

import "time"

//UserProfileIllusts https://www.pixiv.net/ajax/user/<user_id>/profile/illusts?ids[]=<illust_id>&lang=<lang>&version=<version>
type UserProfileIllusts struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Works map[string]struct {
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
				WorkTitle   string `json:"workTitle"`   //WorkTitle 作品标题的翻译
				WorkCaption string `json:"workCaption"` //WorkCaption 作品描述的翻译
			} `json:"titleCaptionTranslation"` //TitleCaptionTranslation 作品标题和描述的翻译
			CreateDate      time.Time `json:"createDate"`
			UpdateDate      time.Time `json:"updateDate"`
			IsUnlisted      bool      `json:"isUnlisted"`
			IsMasked        bool      `json:"isMasked"`
			AiType          int       `json:"aiType"`
			ProfileImageURL string    `json:"profileImageUrl"`
		} `json:"works"`
		ZoneConfig struct {
			Header struct {
				URL string `json:"url"`
			} `json:"header"`
			Footer struct {
				URL string `json:"url"`
			} `json:"footer"`
			Logo struct {
				URL string `json:"url"`
			} `json:"logo"`
			Five00X500 struct {
				URL string `json:"url"`
			} `json:"500x500"`
		} `json:"zoneConfig"`
		ExtraData struct {
			Meta struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				Canonical   string `json:"canonical"`
				Ogp         struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Type        string `json:"type"`
				} `json:"ogp"`
				Twitter struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Card        string `json:"card"`
				} `json:"twitter"`
				AlternateLanguages struct {
					Ja string `json:"ja"`
					En string `json:"en"`
				} `json:"alternateLanguages"`
				DescriptionHeader string `json:"descriptionHeader"`
			} `json:"meta"`
		} `json:"extraData"`
	} `json:"body"`
}
