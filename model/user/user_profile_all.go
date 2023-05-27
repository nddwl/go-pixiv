package user

import "time"

//UserProfileAll https://www.pixiv.net/ajax/user/<user_id>/profile/all?lang=<lang>&version=<version>
type UserProfileAll struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illusts     map[string]struct{} `json:"illust"`
		Manga       map[string]struct{} `json:"manga"`
		Novels      map[string]struct{} `json:"novels"`
		MangaSeries []struct {
			ID             string      `json:"id"`
			UserID         string      `json:"userId"`
			Title          string      `json:"title"`
			Description    string      `json:"description"`
			Caption        string      `json:"caption"`
			Total          int         `json:"total"`
			ContentOrder   interface{} `json:"content_order"`
			URL            string      `json:"url"`
			CoverImageSl   int         `json:"coverImageSl"`
			FirstIllustID  string      `json:"firstIllustId"`
			LatestIllustID string      `json:"latestIllustId"`
			CreateDate     time.Time   `json:"createDate"`
			UpdateDate     time.Time   `json:"updateDate"`
			WatchCount     interface{} `json:"watchCount"`
			IsWatched      bool        `json:"isWatched"`
			IsNotifying    bool        `json:"isNotifying"`
		} `json:"mangaSeries"`
		NovelSeries []struct {
			ID                            string        `json:"id"`
			UserID                        string        `json:"userId"`
			UserName                      string        `json:"userName"`
			ProfileImageURL               string        `json:"profileImageUrl"`
			XRestrict                     int           `json:"xRestrict"`
			IsOriginal                    bool          `json:"isOriginal"`
			IsConcluded                   bool          `json:"isConcluded"`
			GenreID                       string        `json:"genreId"`
			Title                         string        `json:"title"`
			Caption                       string        `json:"caption"`
			Language                      string        `json:"language"`
			Tags                          []interface{} `json:"tags"`
			PublishedContentCount         int           `json:"publishedContentCount"`
			PublishedTotalCharacterCount  int           `json:"publishedTotalCharacterCount"`
			PublishedTotalWordCount       int           `json:"publishedTotalWordCount"`
			PublishedReadingTime          int           `json:"publishedReadingTime"`
			UseWordCount                  bool          `json:"useWordCount"`
			LastPublishedContentTimestamp int           `json:"lastPublishedContentTimestamp"`
			CreatedTimestamp              int           `json:"createdTimestamp"`
			UpdatedTimestamp              int           `json:"updatedTimestamp"`
			CreateDate                    time.Time     `json:"createDate"`
			UpdateDate                    time.Time     `json:"updateDate"`
			FirstNovelID                  string        `json:"firstNovelId"`
			LatestNovelID                 string        `json:"latestNovelId"`
			DisplaySeriesContentCount     int           `json:"displaySeriesContentCount"`
			ShareText                     string        `json:"shareText"`
			Total                         int           `json:"total"`
			FirstEpisode                  struct {
				URL string `json:"url"`
			} `json:"firstEpisode"`
			WatchCount   interface{} `json:"watchCount"`
			MaxXRestrict interface{} `json:"maxXRestrict"`
			Cover        struct {
				Urls struct {
					Two40Mw     string `json:"240mw"`
					Four80Mw    string `json:"480mw"`
					One200X1200 string `json:"1200x1200"`
					One28X128   string `json:"128x128"`
					Original    string `json:"original"`
				} `json:"urls"`
			} `json:"cover"`
			CoverSettingData interface{} `json:"coverSettingData"`
			IsWatched        bool        `json:"isWatched"`
			IsNotifying      bool        `json:"isNotifying"`
			AiType           int         `json:"aiType"`
		} `json:"novelSeries"`
		Pickup []struct {
			Type            string   `json:"type"`
			Deletable       bool     `json:"deletable"`
			Draggable       bool     `json:"draggable"`
			UserName        string   `json:"userName"`
			UserImageURL    string   `json:"userImageUrl"`
			ContentURL      string   `json:"contentUrl"`
			Description     string   `json:"description"`
			ImageURL        string   `json:"imageUrl"`
			ImageURLMobile  string   `json:"imageUrlMobile"`
			HasAdultContent bool     `json:"hasAdultContent"`
			ID              string   `json:"id"`
			Title           string   `json:"title"`
			IllustType      int      `json:"illustType"`
			XRestrict       int      `json:"xRestrict"`
			Restrict        int      `json:"restrict"`
			Sl              int      `json:"sl"`
			URL             string   `json:"url"`
			Tags            []string `json:"tags"`
			UserID          string   `json:"userId"`
			Width           int      `json:"width"`
			Height          int      `json:"height"`
			PageCount       int      `json:"pageCount"`
			IsBookmarkable  bool     `json:"isBookmarkable"`
			BookmarkData    struct {
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
			CreateDate time.Time `json:"createDate"`
			UpdateDate time.Time `json:"updateDate"`
			IsUnlisted bool      `json:"isUnlisted"`
			IsMasked   bool      `json:"isMasked"`
			AiType     int       `json:"aiType"`
			Urls       struct {
				Two50X250   string `json:"250x250"`
				Three60X360 string `json:"360x360"`
				Five40X540  string `json:"540x540"`
			} `json:"urls"`
		} `json:"pickup"`
		BookmarkCount struct {
			Public struct {
				Illust int `json:"illust"`
				Novel  int `json:"novel"`
			} `json:"public"`
			Private struct {
				Illust int `json:"illust"`
				Novel  int `json:"novel"`
			} `json:"private"`
		} `json:"bookmarkCount"`
		ExternalSiteWorksStatus struct {
			Booth    bool `json:"booth"`
			Sketch   bool `json:"sketch"`
			VroidHub bool `json:"vroidHub"`
		} `json:"externalSiteWorksStatus"`
		Request struct {
			ShowRequestTab     bool `json:"showRequestTab"`
			ShowRequestSentTab bool `json:"showRequestSentTab"`
			PostWorks          struct {
				Artworks []string `json:"artworks"` //Artworks 已经请求过的artworks
				Novels   []string `json:"novels"`   //Novels 已经请求过的novels
			} `json:"postWorks"` //PostWorks 已经请求过的数据
		} `json:"request"`
	} `json:"body"`
}
