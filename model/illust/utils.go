package illust

import (
	"bytes"
	"encoding/json"
	"time"
)

//ZoneConfig `json:"zoneConfig"`
type ZoneConfig struct {
	Responsive struct {
		URL string `json:"url"` //URL 链接
	} `json:"responsive"` //Responsive 响应式展示区域的图片URL
	Rectangle struct {
		URL string `json:"url"` //URL 链接
	} `json:"rectangle"` //Rectangle 矩形展示区域的图片URL
	Five00X500 struct {
		URL string `json:"url"` //URL 链接
	} `json:"500x500"` //Five00X500 500 x 500像素展示区域的图片URL
	Header struct {
		URL string `json:"url"` //URL 链接
	} `json:"header"` //Header 页眉展示区域的图片URL
	Footer struct {
		URL string `json:"url"` //URL 链接
	} `json:"footer"` //Footer 页脚展示区域的图片URL
	ExpandedFooter struct {
		URL string `json:"url"` //URL 链接
	} `json:"expandedFooter"` //ExpandedFooter 扩展页脚展示区域的图片URL
	Logo struct {
		URL string `json:"url"` //URL 链接
	} `json:"logo"` //Logo Pixiv标志展示区域的图片URL
	Relatedworks struct {
		URL string `json:"url"` //URL 链接
	} `json:"relatedworks"` //Relatedworks 相关作品展示区域的图片URL
}

//ExtraData `json:"extraData"`
type ExtraData struct {
	Meta struct {
		Title              string `json:"title"`       //Title 标题
		Description        string `json:"description"` //Description 简介
		Canonical          string `json:"canonical"`   //Canonical 规范链接
		AlternateLanguages struct {
			Ja string `json:"ja"` //AlternateLanguages
			En string `json:"en"` //AlternateLanguages
		} `json:"alternateLanguages"` //AlternateLanguages
		DescriptionHeader string `json:"descriptionHeader"` //DescriptionHeader 标题
		Ogp               struct {
			Description string `json:"description"` //Description 协议描述
			Image       string `json:"image"`       //Image 协议图像链接
			Title       string `json:"title"`       //Title 协议标题
			Type        string `json:"type"`        //Type 协议类型
		} `json:"ogp"` //Ogp OpenGraph协议数据
		Twitter struct {
			Description string `json:"description"` //Description Twitter描述
			Image       string `json:"image"`       //Image Twitter图像链接
			Title       string `json:"title"`       //Title Twitter标题
			Card        string `json:"card"`        //Card illust在Twitter上展示方式
		} `json:"twitter"` //Twitter 作者推特
	} `json:"meta"`
}

type ResponseData struct {
	ImageResponseOutData []struct {
		Type     string `json:"type"`     //Type 作品类型[illust,manga,ugoira]
		WorkID   string `json:"workId"`   //WorkID 图片ID
		Title    string `json:"title"`    //Title 图片标题
		UserName string `json:"userName"` //UserName 作者的用户名
		ImageURL string `json:"imageUrl"` //ImageURL 图片的URL地址
	} `json:"ImageResponseOutData"` // 作品是某作品的关联作品
	ImageResponseData []struct {
		Type           string   `json:"type"`           //Type 关联作品类型[illust,manga,ugoira]
		ID             string   `json:"id"`             //ID 关联作品的ID
		Title          string   `json:"title"`          //Title 关联作品的标题
		IllustType     int      `json:"illustType"`     //IllustType 关联作品的类型[0:"illust",1:"manga",2:"ugoira"]
		XRestrict      int      `json:"xRestrict"`      //XRestrict 关联作品的限制级别[0:"未受限制",1:"已受到某些限制,可能需要进行年龄验证或仅限于特定地区或用户",2:"作品已被完全限制,只能在特定条件下查看"]
		Restrict       int      `json:"restrict"`       //Restrict 关联作品的受限状态[0:"public",1:"private",2:"friends",3:"me"]
		Sl             int      `json:"sl"`             //Sl 关联展示级别[0:"普通作品",1:"R-18作品,但不包含性行为",2:"性行为作品",3:"血腥暴力作品",4:"非人类/怪物作品",5:"极度暴力/死亡/酷刑作品",6:"极度变态/性侵作品"]
		URL            string   `json:"url"`            //URL 关联作品地址
		Description    string   `json:"description"`    //Description 关联作品描述
		Tags           []string `json:"tags"`           //Tags 关联作品标签
		UserID         string   `json:"userId"`         //UserID 关联作品作者ID
		UserName       string   `json:"userName"`       //UserName 关联作品作者姓名
		Width          int      `json:"width"`          //Width 关联作品宽度
		Height         int      `json:"height"`         //Height 关联作品高度
		PageCount      int      `json:"pageCount"`      //PageCount 关联作品页数
		IsBookmarkable bool     `json:"isBookmarkable"` //IsBookmarkable 关联作品是否已收藏
		BookmarkData   struct {
			ID      string `json:"id"`      //ID 该收藏记录的唯一标识符
			Private bool   `json:"private"` //Private 是否为私人收藏
		} `json:"bookmarkData"` //BookmarkData 收藏数据
		Alt                     string `json:"alt"` //Alt 关联作品的替代描述
		TitleCaptionTranslation struct {
			WorkTitle   string `json:"workTitle"`   //WorkTitle 关联作品标题的翻译
			WorkCaption string `json:"workCaption"` //WorkCaption 关联作品描述的翻译
		} `json:"titleCaptionTranslation"` //TitleCaptionTranslation 关联作品标题和描述的翻译
		CreateDate time.Time `json:"createDate"` //CreateDate 关联作品上传日期
		UpdateDate time.Time `json:"updateDate"` //UpdateDate 关联作品更新日期
		IsUnlisted bool      `json:"isUnlisted"` //IsUnlisted 关联作品是否被设为未公开状态
		IsMasked   bool      `json:"isMasked"`   //IsMasked 关联作品是否被屏蔽
		AiType     int       `json:"aiType"`     //AiType 关联作品是否为AI作品
		Urls       struct {
			Two50X250   string `json:"250x250"` //Two50X250 尺寸
			Three60X360 string `json:"360x360"` //Three60X360 尺寸
			Five40X540  string `json:"540x540"` //Five40X540 尺寸
		} `json:"urls"` //Urls 关联作品图片链接
		ProfileImageURL string `json:"profileImageUrl"` //ProfileImageURL 关联作品的作者头像链接
	} `json:"imageResponseData"` // 作品的关联作品信息
	ImageResponseCount int `json:"imageResponseCount"` //ImageResponseCount 关联作品的数量
}

//SeriesNavData  `json:"seriesNavData"` 作品是否属于一个系列
type SeriesNavData struct {
	SeriesType  string `json:"seriesType"`  //SeriesType 系列类型,表示这个作品是否属于一个系列取值为字符串类型
	SeriesID    string `json:"seriesId"`    //SeriesID 系列 ID,如果这个作品属于一个系列,那么这个字段将会是系列的 ID取值为字符串类型
	Title       string `json:"title"`       //Title 系列标题,表示系列的名称取值为字符串类型
	Order       int    `json:"order"`       //Order 系列顺序,表示这个作品在系列中的顺序取值为整型
	IsWatched   bool   `json:"isWatched"`   //IsWatched 表示用户是否观看过这个作品
	IsNotifying bool   `json:"isNotifying"` //IsNotifying 表示用户是否开启了该系列的通知功能
	Prev        struct {
		ID    string `json:"id"`    //ID 前一个作品的 ID取值为字符串类型
		Title string `json:"title"` //Title 前一个作品的标题取值为字符串类型
		Order int    `json:"order"` //Order 前一个作品在系列中的顺序取值为整型
	} `json:"prev"` //Prev 前一个作品,表示系列中上一个作品的信息
	Next struct {
		ID    string `json:"id"`    //ID 后一个作品的 ID取值为字符串类型
		Title string `json:"title"` //Title 后一个作品的标题取值为字符串类型
		Order int    `json:"order"` //Order 后一个作品在系列中的顺序取值为整型
	} `json:"next"` //Next 后一个作品,表示系列中下一个作品的信息
}

//FanboxPromotion `json:"fanboxPromotion"` Fanbox信息
type FanboxPromotion struct {
	UserName        string `json:"userName"`        //UserName 作者名称
	UserImageURL    string `json:"userImageUrl"`    //UserImageURL 作者的头像url
	ContentURL      string `json:"contentUrl"`      //ContentURL Fanbox的主页URL
	Description     string `json:"description"`     //Description Fanbox的描述
	ImageURL        string `json:"imageUrl"`        //ImageURL Fanbox的封面图片URL
	ImageURLMobile  string `json:"imageUrlMobile"`  //ImageURLMobile Fanbox的移动版封面图片URL
	HasAdultContent bool   `json:"hasAdultContent"` //HasAdultContent Fanbox是否含有成人内容
}

//ContestBanners `json:"contestBanners"` pixiv竞赛横幅相关信息
type ContestBanners struct {
	IllustID                            string    `json:"illust_id"`                               //IllustID 插画的唯一标识符
	IllustUserID                        string    `json:"illust_user_id"`                          //IllustUserID 插画的作者 ID
	IllustTitle                         string    `json:"illust_title"`                            //IllustTitle 插画的标题
	IllustExt                           string    `json:"illust_ext"`                              //IllustExt 插画文件的扩展名
	IllustWidth                         string    `json:"illust_width"`                            //IllustWidth 插画的宽度
	IllustHeight                        string    `json:"illust_height"`                           //IllustHeight 插画的高度
	IllustRestrict                      string    `json:"illust_restrict"`                         //IllustRestrict 插画的限制级别
	IllustXRestrict                     string    `json:"illust_x_restrict"`                       //IllustXRestrict 插画的 R18 级别
	IllustCreateDate                    string    `json:"illust_create_date"`                      //IllustCreateDate 插画的上传日期
	IllustUploadDate                    string    `json:"illust_upload_date"`                      //IllustUploadDate 插画的更新日期
	IllustServerID                      string    `json:"illust_server_id"`                        //IllustServerID 插画存储的服务器 ID
	IllustHash                          string    `json:"illust_hash"`                             //IllustHash 插画的哈希值
	IllustType                          string    `json:"illust_type"`                             //IllustType 插画的类型
	IllustSanityLevel                   int       `json:"illust_sanity_level"`                     //IllustSanityLevel 插画的合理性等级
	IllustBookStyle                     int       `json:"illust_book_style"`                       //IllustBookStyle 插画是否为书籍样式
	IllustPageCount                     string    `json:"illust_page_count"`                       //IllustPageCount 插画的页数
	IllustCommentOffSetting             string    `json:"illust_comment_off_setting"`              //IllustCommentOffSetting 插画的评论关闭设置
	IllustAiType                        string    `json:"illust_ai_type"`                          //IllustAiType 插画的 AI 类型
	IllustCustomThumbnailUploadDatetime time.Time `json:"illust_custom_thumbnail_upload_datetime"` //IllustCustomThumbnailUploadDatetime 自定义缩略图上传日期
	IllustComment                       string    `json:"illust_comment"`                          //IllustComment 插画的评论
	IllustTagFullLock                   string    `json:"illust_tag_full_lock"`                    //IllustTagFullLock 插画的标签完全锁定设置
	IllustTool01                        string    `json:"illust_tool01"`                           //IllustTool01 插画的工具 1
	IllustTool02                        string    `json:"illust_tool02"`                           //IllustTool02 插画的工具 2
	IllustTool03                        string    `json:"illust_tool03"`                           //IllustTool03 插画的工具 3
	UserAccount                         string    `json:"user_account"`                            //UserAccount 作者的账户名
	UserName                            string    `json:"user_name"`                               //UserName 作者的用户名
	UserPremium                         string    `json:"user_premium"`                            //UserPremium 作者的会员类型
	URL                                 struct {
		Four8X48         string `json:"48x48"`          //Four8X48 尺寸
		Two50X250        string `json:"250x250"`        //Two50X250 尺寸
		Five40X540Master string `json:"540x540_master"` //Five40X540Master 尺寸
		One200X1200      string `json:"1200x1200"`      //One200X1200 尺寸
		Big              string `json:"big"`            //Big 尺寸
	} `json:"url"` //URL 插画文件的 URL,包括不同尺寸的图片
	IllustRatingCount string `json:"illust_rating_count"` //IllustRatingCount 插画的评价数
	IllustRatingScore string `json:"illust_rating_score"` //IllustRatingScore 插画的评分
	IllustRatingView  string `json:"illust_rating_view"`  //IllustRatingView 插画的浏览数
	IllustContentType struct {
		Sexual     int  `json:"sexual"`     //Sexual 是否包含性内容
		Lo         bool `json:"lo"`         //Lo 是否包含低俗内容
		Grotesque  bool `json:"grotesque"`  //Grotesque 是否包含恶心内容
		Violent    bool `json:"violent"`    //Violent 是否包含暴力内容
		Homosexual bool `json:"homosexual"` //Homosexual 是否包含同性恋内容
		Drug       bool `json:"drug"`       //Drug 是否包含毒品内容
		Thoughts   bool `json:"thoughts"`   //Thoughts 是否包含思考题材内容
		Antisocial bool `json:"antisocial"` //Antisocial 是否包含反社会内容
		Religion   bool `json:"religion"`   //Religion 是否包含宗教内容
		Original   bool `json:"original"`   //Original 是否是原创内容
		Furry      bool `json:"furry"`      //Furry 是否包含福瑞属性
		Bl         bool `json:"bl"`         //Bl 是否是男同性恋内容
		Yuri       bool `json:"yuri"`       //Yuri 是否是女同性恋内容
	} `json:"illust_content_type"` //IllustContentType 插画的内容类型,包括性爱 低俗 恶心 暴力 同性恋 毒品 反社会 宗教 原创 福瑞等
	Tags         []string `json:"tags"`          //Tags 插画的标签
	IllustSeries bool     `json:"illust_series"` //IllustSeries 插画是否属于系列
	IsBookmarked bool     `json:"is_bookmarked"` //IsBookmarked 当前用户是否已将此插画收藏
	Bookmarkable bool     `json:"bookmarkable"`  //Bookmarkable 插画是否可以被收藏
	Res          int      `json:"res"`           //Res 分辨率
	Cnt          int      `json:"cnt"`           //Cnt 计数
	DisplayTags  []struct {
		Tag                     string `json:"tag"`                        //Tag 标签名称
		IsMyself                bool   `json:"is_myself"`                  //IsMyself 是否为自己添加的标签
		HasLock                 bool   `json:"has_lock"`                   //HasLock 是否被锁定
		AddUserID               int    `json:"add_user_id"`                //AddUserID 添加标签的用户
		NameFlg                 bool   `json:"name_flg"`                   //NameFlg 账户类型[0:"普通用户",1:"官方账户",2:"画师账户",3:"小号",4:"被冻结的账户"]
		IsPixpediaArticleExists bool   `json:"is_pixpedia_article_exists"` //IsPixpediaArticleExists Pixpedia上是否存在指定标题的文章(Pixpedia是Pixiv推出的一个百科全书式的内容分享平台,用户可以在上面创建和编辑关于动漫 漫画 游戏 小说等相关话题的条目)
		Translation             string `json:"translation"`                //Translation 翻译(可能与lang=zn;en有关)
	} `json:"display_tags"` //DisplayTags 展示的标签,包括标签名称 是否为自己添加的标签 是否被锁定 添加标签的用户ID 账户类型 是否为Pixpedia文章 翻译
}

type Detail struct {
	Methods         []string `json:"methods"`         //Methods 表示用于推荐该作品的方法或算法的列表
	Score           float64  `json:"score"`           //Score 表示该作品的推荐分数
	SeedIllustIds   []string `json:"seedIllustIds"`   //SeedIllustIds 表示作为种子作品用于推荐该作品的其他作品的 ID 列表
	BanditInfo      string   `json:"banditInfo"`      //BanditInfo 表示与推荐算法相关的带宽信息
	RecommendListID string   `json:"recommendListId"` //RecommendListID 表示推荐列表的唯一标识符
}

type Details struct {
	Detail []Detail
}

func (t *Details) UnmarshalJSON(data []byte) error {
	if t == nil {
		*t = Details{}
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.Token()
	for decoder.More() {
		var detail Detail
		decoder.Token()
		decoder.Decode(&detail)
		t.Detail = append(t.Detail, detail)
	}
	return nil
}
