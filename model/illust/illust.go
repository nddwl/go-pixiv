package illust

import (
	"go-pixiv/model/utils"
	"time"
)

//Illust https://www.pixiv.net/ajax/illust/<illust_id>?lang=<lang>&version=<version>
type Illust struct {
	Error   bool   `json:"error"`   // Error 表示是否有错误,如果没有错误则为false,有错误则为true
	Message string `json:"message"` //Message 错误信息
	Body    struct {
		IllustID      string    `json:"illustId"`      //IllustID 作品ID
		IllustTitle   string    `json:"illustTitle"`   //IllustTitle 作品标题
		IllustComment string    `json:"illustComment"` //IllustComment 作品描述
		ID            string    `json:"id"`            //ID 作品ID
		Title         string    `json:"title"`         //Title 作品标题
		Description   string    `json:"description"`   //Description 作品描述
		IllustType    int       `json:"illustType"`    //IllustType 作品类型[0:"illust",1:"manga",2:"ugoira"]
		CreateDate    time.Time `json:"createDate"`    //CreateDate 上传时间
		UploadDate    time.Time `json:"uploadDate"`    //UploadDate 最后更新时间
		Restrict      int       `json:"restrict"`      //Restrict 作品公开状态[0:"public",1:"private",2:"friends",3:"me"]
		XRestrict     int       `json:"xRestrict"`     //XRestrict 作品限制级别[0:"未受限制",1:"已受到某些限制,可能需要进行年龄验证或仅限于特定地区或用户",2:"作品已被完全限制,只能在特定条件下查看"]
		Sl            int       `json:"sl"`            //Sl 展示级别[0:"普通作品",1:"R-18作品,但不包含性行为",2:"性行为作品",3:"血腥暴力作品",4:"非人类/怪物作品",5:"极度暴力/死亡/酷刑作品",6:"极度变态/性侵作品"]
		Urls          struct {
			Mini     string `json:"mini"`
			Thumb    string `json:"thumb"`
			Small    string `json:"small"`
			Regular  string `json:"regular"`
			Original string `json:"original"`
		} `json:"urls"` //Urls 作品链接
		Tags          utils.Tags   `json:"tags"`          //Tags 作品标签
		Alt           string       `json:"alt"`           //Alt 作品替代文本
		StorableTags  []string     `json:"storableTags"`  //StorableTags 用户收藏作品时可以存储的标签列表
		UserID        string       `json:"userId"`        //UserID 上传用户ID
		UserName      string       `json:"userName"`      //UserName 上传用户名称
		UserAccount   string       `json:"userAccount"`   //UserAccount 上传用户账户名称
		UserIllusts   utils.Illust `json:"userIllusts"`   //UserIllusts 用户作品列表
		LikeData      bool         `json:"likeData"`      //LikeData 是否已点赞
		Width         int          `json:"width"`         //Width 作品宽度
		Height        int          `json:"height"`        //Height 作品高度
		PageCount     int          `json:"pageCount"`     //PageCount 作品页数
		BookmarkCount int          `json:"bookmarkCount"` //BookmarkCount 作品被收藏次数
		LikeCount     int          `json:"likeCount"`     //LikeCount 作品被点赞次数
		CommentCount  int          `json:"commentCount"`  //CommentCount 作品被评论次数
		ResponseCount int          `json:"responseCount"` //ResponseCount 作品被回复次数
		ViewCount     int          `json:"viewCount"`     //ViewCount 作品被浏览次数
		BookStyle     string       `json:"bookStyle"`     //BookStyle 作品画集类型
		IsHowto       bool         `json:"isHowto"`       //IsHowto 作品是否为教程/指南类作品
		IsOriginal    bool         `json:"isOriginal"`    //IsOriginal 作品是否为原创作品的,二次创作作品,合作作品等
		ResponseData
		PollData             interface{}     `json:"pollData"`             //PollData 投票信息
		SeriesNavData        SeriesNavData   `json:"seriesNavData"`        //SeriesNavData 作品是否属于一个系列
		DescriptionBoothID   int             `json:"descriptionBoothId"`   //DescriptionBoothID BoothID的展位ID
		DescriptionYoutubeID int             `json:"descriptionYoutubeId"` //DescriptionYoutubeID Youtube的展位ID
		ComicPromotion       interface{}     `json:"comicPromotion"`       //ComicPromotion 作品竞赛信息
		FanboxPromotion      FanboxPromotion `json:"fanboxPromotion"`      //FanboxPromotion Fanbox信息
		ContestBanners       ContestBanners  `json:"contestBanners"`       //ContestBanners pixiv竞赛横幅相关信息
		IsBookmarkable       bool            `json:"isBookmarkable"`       //IsBookmarkable 是否收藏
		BookmarkData         struct {
			ID      string `json:"id"`      //ID 作品书签ID
			Private bool   `json:"private"` //Private 是否为私人书签
		} `json:"bookmarkData"` //BookmarkData 书签数据
		ContestData struct {
			URL   string `json:"url"`   //URL 比赛链接地址
			Icon  string `json:"icon"`  //Icon 比赛内容图标
			Title string `json:"title"` //Title 比赛内容标题
		} `json:"contestData"` //ContestData 比赛有关数据
		ZoneConfig              ZoneConfig `json:"zoneConfig"` //ZoneConfig 绘画作品在不同展示区域,比如首页,排行榜,搜索结果页面等显示时所使用的图片资源的URL配置
		ExtraData               ExtraData  `json:"extraData"`  //ExtraData 额外数据
		TitleCaptionTranslation struct {
			WorkTitle struct {
				En string
				Ja string
			} `json:"workTitle"` //WorkTitle 标题翻译
			WorkCaption struct {
				En string
				Ja string
			} `json:"workCaption"` //WorkCaption 描述翻译
		} `json:"titleCaptionTranslation"` //作品标题和作品描述的翻译信息
		IsUnlisted bool        `json:"isUnlisted"` //IsUnlisted 是否被设为不公开的
		Request    interface{} `json:"request"`    //Request 管理illust的请求
		CommentOff int         `json:"commentOff"` //CommentOff 评论是否关闭
		AiType     int         `json:"aiType"`     //AiType AI分级[0:非AI作品,1:色情图片识别模型,2:草稿模型,3:扫描动漫和插图的模型,4:扫描摄影和背景的模型,5:扫描普通的摄影和插图的模型,6:扫描儿童插图的模型]
	} `json:"body"`
}
