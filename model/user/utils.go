package user

import (
	"bytes"
	"encoding/json"
	"time"
)

type Manga struct {
	ID             string      `json:"id"`             // 漫画系列ID
	UserID         string      `json:"userId"`         // 用户ID
	Title          string      `json:"title"`          // 漫画标题
	Description    string      `json:"description"`    // 漫画描述
	Caption        string      `json:"caption"`        // 漫画副标题
	Total          int         `json:"total"`          // 漫画总页数
	ContentOrder   interface{} `json:"content_order"`  // 漫画内容顺序
	URL            string      `json:"url"`            // 漫画URL
	CoverImageSl   int         `json:"coverImageSl"`   // 封面图片的Sl值
	FirstIllustID  string      `json:"firstIllustId"`  // 第一张插画的ID
	LatestIllustID string      `json:"latestIllustId"` // 最新一张插画的ID
	CreateDate     time.Time   `json:"createDate"`     // 创建日期
	UpdateDate     time.Time   `json:"updateDate"`     // 更新日期
	WatchCount     interface{} `json:"watchCount"`     // 观看计数
	IsWatched      bool        `json:"isWatched"`      // 是否已观看
	IsNotifying    bool        `json:"isNotifying"`    // 是否接收通知
}

//MangaSeries [] json:"mangaSeries"
type MangaSeries struct {
	Manga []Manga
}

func (t *MangaSeries) UnmarshalJSON(data []byte) error {
	if t == nil {
		*t = MangaSeries{}
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.Token()
	for decoder.More() {
		var manga Manga
		decoder.Token()
		decoder.Decode(&manga)
		t.Manga = append(t.Manga, manga)
	}
	return nil
}

type Novel struct {
	ID                            string        `json:"id"`                            // 小说系列的唯一ID
	UserID                        string        `json:"userId"`                        // 用户的唯一ID
	UserName                      string        `json:"userName"`                      // 用户名
	ProfileImageURL               string        `json:"profileImageUrl"`               // 用户的个人资料图片URL
	XRestrict                     int           `json:"xRestrict"`                     // 内容限制级别
	IsOriginal                    bool          `json:"isOriginal"`                    // 是否为原创系列
	IsConcluded                   bool          `json:"isConcluded"`                   // 是否已完结
	GenreID                       string        `json:"genreId"`                       // 小说系列所属的流派ID
	Title                         string        `json:"title"`                         // 小说系列的标题
	Caption                       string        `json:"caption"`                       // 小说系列的描述
	Language                      string        `json:"language"`                      // 小说系列的语言
	Tags                          []interface{} `json:"tags"`                          // 小说系列的标签
	PublishedContentCount         int           `json:"publishedContentCount"`         // 发布的内容数量
	PublishedTotalCharacterCount  int           `json:"publishedTotalCharacterCount"`  // 发布的内容总字符数
	PublishedTotalWordCount       int           `json:"publishedTotalWordCount"`       // 发布的内容总字数
	PublishedReadingTime          int           `json:"publishedReadingTime"`          // 发布的内容总阅读时间
	UseWordCount                  bool          `json:"useWordCount"`                  // 是否使用字数统计
	LastPublishedContentTimestamp int           `json:"lastPublishedContentTimestamp"` // 最后发布内容的时间戳
	CreatedTimestamp              int           `json:"createdTimestamp"`              // 创建时间戳
	UpdatedTimestamp              int           `json:"updatedTimestamp"`              // 更新时间戳
	CreateDate                    time.Time     `json:"createDate"`                    // 创建日期
	UpdateDate                    time.Time     `json:"updateDate"`                    // 更新日期
	FirstNovelID                  string        `json:"firstNovelId"`                  // 第一篇小说的ID
	LatestNovelID                 string        `json:"latestNovelId"`                 // 最新一篇小说的ID
	DisplaySeriesContentCount     int           `json:"displaySeriesContentCount"`     // 显示的系列内容数量
	ShareText                     string        `json:"shareText"`                     // 分享文本
	Total                         int           `json:"total"`                         // 总计数
	FirstEpisode                  struct {
		URL string `json:"url"` // 第一篇小说的URL
	} `json:"firstEpisode"`
	WatchCount   interface{} `json:"watchCount"`   // 观看计数
	MaxXRestrict interface{} `json:"maxXRestrict"` // 最大限制级别
	Cover        struct {
		Urls struct {
			Two40Mw     string `json:"240mw"`     // 封面图片 URL (240mw)
			Four80Mw    string `json:"480mw"`     // 封面图片 URL (480mw)
			One200X1200 string `json:"1200x1200"` // 封面图片 URL (1200x1200)
			One28X128   string `json:"128x128"`   // 封面图片 URL (128x128)
			Original    string `json:"original"`  // 封面图片 URL (原始尺寸)
		} `json:"urls"`
	} `json:"cover"`
	CoverSettingData interface{} `json:"coverSettingData"` // 封面设置数据
	IsWatched        bool        `json:"isWatched"`        // 是否已观看
	IsNotifying      bool        `json:"isNotifying"`      // 是否接收通知
	AiType           int         `json:"aiType"`           // AI类型
}

//NovelSeries [] json:"novelSeries"
type NovelSeries struct {
	Novel []Novel
}

func (t *NovelSeries) UnmarshalJSON(data []byte) error {
	if t == nil {
		*t = NovelSeries{}
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.Token()
	for decoder.More() {
		var novel Novel
		decoder.Token()
		decoder.Decode(&novel)
		t.Novel = append(t.Novel, novel)
	}
	return nil
}

type PickupWork struct {
	Type            string   `json:"type"`            // 类型
	Deletable       bool     `json:"deletable"`       // 是否可删除
	Draggable       bool     `json:"draggable"`       // 是否可拖动
	UserName        string   `json:"userName"`        // 用户名
	UserImageURL    string   `json:"userImageUrl"`    // 用户头像 URL
	ContentURL      string   `json:"contentUrl"`      // 内容 URL
	Description     string   `json:"description"`     // 描述
	ImageURL        string   `json:"imageUrl"`        // 图片 URL
	ImageURLMobile  string   `json:"imageUrlMobile"`  // 移动端图片 URL
	HasAdultContent bool     `json:"hasAdultContent"` // 是否包含成人内容
	ID              string   `json:"id"`              // ID
	Title           string   `json:"title"`           // 标题
	IllustType      int      `json:"illustType"`      // 作品类型
	XRestrict       int      `json:"xRestrict"`       // xRestrict
	Restrict        int      `json:"restrict"`        // 限制
	Sl              int      `json:"sl"`              // sl
	URL             string   `json:"url"`             // URL
	Tags            []string `json:"tags"`            // 标签
	UserID          string   `json:"userId"`          // 用户 ID
	Width           int      `json:"width"`           // 宽度
	Height          int      `json:"height"`          // 高度
	PageCount       int      `json:"pageCount"`       // 页数
	IsBookmarkable  bool     `json:"isBookmarkable"`  // 是否可加入书签
	BookmarkData    struct {
		ID      string `json:"id"`      // 作品唯一书签 ID
		Private bool   `json:"private"` // 是否为私人书签
	} `json:"bookmarkData"` // 书签数据
	Alt                     string `json:"alt"` // Alt
	TitleCaptionTranslation struct {
		WorkTitle struct {
			En string `json:"en"` // 英文标题
			Ja string `json:"ja"` // 日文标题（默认）
		} `json:"workTitle"` // 作品标题翻译
		WorkCaption struct {
			En string `json:"en"` // 英文描述
			Ja string `json:"ja"` // 日文描述（默认）
		} `json:"workCaption"` // 作品描述翻译
	} `json:"titleCaptionTranslation"` // 作品标题和描述的翻译信息
	CreateDate time.Time `json:"createDate"` // 创建日期
	UpdateDate time.Time `json:"updateDate"` // 更新日期
	IsUnlisted bool      `json:"isUnlisted"` // 是否不公开
	IsMasked   bool      `json:"isMasked"`   // 是否遮罩
	AiType     int       `json:"aiType"`     // AI 类型
	Urls       struct {
		Two50X250   string `json:"250x250"` // 250x250 图片 URL
		Three60X360 string `json:"360x360"` // 360x360 图片 URL
		Five40X540  string `json:"540x540"` // 540x540 图片 URL
	} `json:"urls"` // 图片 URL
}

//Pickup [] json:"pickup"
type Pickup struct {
	PickupWork []PickupWork
}

func (t *Pickup) UnmarshalJSON(data []byte) error {
	if t == nil {
		*t = Pickup{}
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.Token()
	for decoder.More() {
		var work PickupWork
		decoder.Token()
		decoder.Decode(&work)
		t.PickupWork = append(t.PickupWork, work)
	}
	return nil
}
