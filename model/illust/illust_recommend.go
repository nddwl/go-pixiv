package illust

import "time"

//IllustRecommend https://www.pixiv.net/ajax/illust/<illust_id>/recommend/init?limit=<num>&lang=<lang>&version=<version>
// limit控制返回illusts[]数量
type IllustRecommend struct {
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
			ProfileImageURL string    `json:"profileImageUrl"` //ProfileImageURL 作者头像链接
			Type            string    `json:"type"`
			IsAdContainer   bool      `json:"isAdContainer"` //IsAdContainer 是否为广告
		} `json:"illust"`
		NextIds []string `json:"nextIds"` //NextIds 表示下一批推荐作品的 ID 列表
		Details map[string]struct {
			Methods         []string `json:"methods"`         //Methods 表示用于推荐该作品的方法或算法的列表
			Score           float64  `json:"score"`           //Score 表示该作品的推荐分数
			SeedIllustIds   []string `json:"seedIllustIds"`   //SeedIllustIds 表示作为种子作品用于推荐该作品的其他作品的 ID 列表
			BanditInfo      string   `json:"banditInfo"`      //BanditInfo 表示与推荐算法相关的带宽信息
			RecommendListID string   `json:"recommendListId"` //RecommendListID 表示推荐列表的唯一标识符
		} `json:"details"` //Details illusts中作品的详细信息
	} `json:"body"`
}
