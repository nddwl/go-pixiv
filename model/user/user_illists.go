package user

import "time"

//UserIllusts https://www.pixiv.net/ajax/user/<user_id>/illusts?ids[]=<illust_id>&lang=<lang>&version=<version>
type UserIllusts struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    map[string]struct {
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
			ID      string `json:"id"`      //ID 作品唯一书签ID
			Private bool   `json:"private"` //Private 是否为私人书签
		} `json:"bookmarkData"` //BookmarkData 这是有关书签数据的信息,包括是否已加入书签,收藏次数等
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
	} `json:"body"`
}
