package user

import "time"

//UserWorksLatest https://www.pixiv.net/ajax/user/<user_id>/works/latest?lang=<lang>&version=<version>  只有前几个有值
type UserWorksLatest struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illusts map[string]struct {
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
		} `json:"illusts"`
		Novels map[string]struct {
			ID              string   `json:"id"`
			Title           string   `json:"title"`
			XRestrict       int      `json:"xRestrict"`
			Restrict        int      `json:"restrict"`
			URL             string   `json:"url"`
			Tags            []string `json:"tags"`
			UserID          string   `json:"userId"`
			UserName        string   `json:"userName"`
			ProfileImageURL string   `json:"profileImageUrl"`
			TextCount       int      `json:"textCount"`
			WordCount       int      `json:"wordCount"`
			ReadingTime     int      `json:"readingTime"`
			UseWordCount    bool     `json:"useWordCount"`
			Description     string   `json:"description"`
			IsBookmarkable  bool     `json:"isBookmarkable"`
			BookmarkData    struct {
				ID      string `json:"id"`      //ID 该收藏记录的唯一标识符,可以用于后续的修改或删除操作
				Private bool   `json:"private"` //Private 是否为私人收藏,如果为 true,则表示这个收藏仅对用户自己可见,如果为 false,则表示这个收藏为公开的
			} `json:"bookmarkData"` //BookmarkData 收藏数据
			BookmarkCount           int         `json:"bookmarkCount"`
			IsOriginal              bool        `json:"isOriginal"`
			Marker                  interface{} `json:"marker"`
			TitleCaptionTranslation struct {
				WorkTitle   string `json:"workTitle"`   //WorkTitle 作品标题的翻译
				WorkCaption string `json:"workCaption"` //WorkCaption 作品描述的翻译
			} `json:"titleCaptionTranslation"` //TitleCaptionTranslation 作品标题和描述的翻译
			CreateDate time.Time `json:"createDate"`
			UpdateDate time.Time `json:"updateDate"`
			IsMasked   bool      `json:"isMasked"`
			IsUnlisted bool      `json:"isUnlisted"`
			AiType     int       `json:"aiType"`
		} `json:"novels"`
	} `json:"body"`
}
