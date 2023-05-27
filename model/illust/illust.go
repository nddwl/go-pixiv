package illust

import "time"

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
		Restrict      int       `json:"restrict"`      //Restrict 表示作品的公开状态[0:"public",1:"private",2:"friends",3:"me"]
		XRestrict     int       `json:"xRestrict"`     //XRestrict 作品的限制级别[0:"未受限制",1:"已受到某些限制,可能需要进行年龄验证或仅限于特定地区或用户",2:"作品已被完全限制,只能在特定条件下查看"]
		Sl            int       `json:"sl"`            //Sl 展示级别[0:"普通作品",1:"R-18作品,但不包含性行为",2:"性行为作品",3:"血腥暴力作品",4:"非人类/怪物作品",5:"极度暴力/死亡/酷刑作品",6:"极度变态/性侵作品"]
		Urls          struct {
			Mini     string `json:"mini"`     //Mini 图片的小尺寸缩略图链接
			Thumb    string `json:"thumb"`    //Thumb 图片的较小尺寸缩略图链接
			Small    string `json:"small"`    //Small 图片的中等尺寸缩略图链接
			Regular  string `json:"regular"`  //Regular 图片的常规尺寸链接
			Original string `json:"original"` //Original 图片的原始尺寸链接,即最高分辨率的图片链接
		} `json:"urls"` //Urls 作品链接
		Tags struct {
			AuthorID string `json:"authorId"` //AuthorID 作者id
			IsLocked bool   `json:"isLocked"` //IsLocked 是否被锁定,如果为true,则无法编辑或删除这个标签
			Tags     []struct {
				Tag         string `json:"tag"`       //Tag 标签名称
				Locked      bool   `json:"locked"`    //Locked 是否被锁定,如果为true,则无法编辑或删除这个标签
				Deletable   bool   `json:"deletable"` //Deletable 是否可删除,如果为true,则有权删除此标签
				UserID      string `json:"userId"`    //UserID 该标签的创建者用户ID
				UserName    string `json:"userName"`  //UserName 该标签的创建者用户名
				Translation struct {
					En string `json:"en"` //En 标签的根据用户的语言或lang来返回结果,具体会根据http请求中lang值发送变化(例lang=zn,则这里为中文)
				} `json:"translation"` //Translation 标签的翻译对象
			} `json:"tags"` //Tags 具体标签
			Writable bool `json:"writable"` //Writable 是否可以编辑这个作品的标签,如果为true,则有权添加或删除这个作品的标签
		} `json:"tags"` //Tags 作品标签
		Alt          string   `json:"alt"`          //Alt 作品标题或描述的替代文本
		StorableTags []string `json:"storableTags"` //StorableTags 用户收藏作品时可以存储的标签列表
		UserID       string   `json:"userId"`       //UserID 上传作品的用户的ID
		UserName     string   `json:"userName"`     //UserName 上传作品的用户的ID
		UserAccount  string   `json:"userAccount"`  //UserAccount 上传作品的用户的账户名称
		UserIllusts  map[string]struct {
			ID             string   `json:"id"`             //ID 作品ID
			Title          string   `json:"title"`          //Title 作品标题
			IllustType     int      `json:"illustType"`     //IllustType 作品类型[0:"illust",1:"manga",2:"ugoira"]
			XRestrict      int      `json:"xRestrict"`      //XRestrict 作品的限制级别[0:"未受限制",1:"已受到某些限制,可能需要进行年龄验证或仅限于特定地区或用户",2:"作品已被完全限制,只能在特定条件下查看"]
			Restrict       int      `json:"restrict"`       //Restrict 表示作品的受限状态[0:"public",1:"private",2:"friends",3:"me"]
			Sl             int      `json:"sl"`             //Sl 展示级别[0:"普通作品",1:"R-18作品,但不包含性行为",2:"性行为作品",3:"血腥暴力作品",4:"非人类/怪物作品",5:"极度暴力/死亡/酷刑作品",6:"极度变态/性侵作品"]
			URL            string   `json:"url"`            //URL 作品链接
			Description    string   `json:"description"`    //Description 作品描述
			Tags           []string `json:"tags"`           //Tags 作品标签
			UserID         string   `json:"userId"`         //UserID 用户ID
			UserName       string   `json:"userName"`       //UserName 用户名
			Width          int      `json:"width"`          //Width 作品宽度
			Height         int      `json:"height"`         //Height 作品高度
			PageCount      int      `json:"pageCount"`      //PageCount 作品页数
			IsBookmarkable bool     `json:"isBookmarkable"` //IsBookmarkable 是否已收藏
			BookmarkData   struct {
				ID      string `json:"id"`      //ID 该收藏记录的唯一标识符,可以用于后续的修改或删除操作
				Private bool   `json:"private"` //Private 是否为私人收藏,如果为 true,则表示这个收藏仅对用户自己可见,如果为 false,则表示这个收藏为公开的
			} `json:"bookmarkData"` //BookmarkData 收藏数据
			Alt                     string `json:"alt"` //Alt 作品标题或描述的替代文本
			TitleCaptionTranslation struct {
				WorkTitle   string `json:"workTitle"`   //WorkTitle 作品标题的翻译
				WorkCaption string `json:"workCaption"` //WorkCaption 作品描述的翻译
			} `json:"titleCaptionTranslation"` //TitleCaptionTranslation 作品标题和描述的翻译
			CreateDate      time.Time `json:"createDate"`      //CreateDate 上传时间
			UpdateDate      time.Time `json:"updateDate"`      //UpdateDate 更新时间
			IsUnlisted      bool      `json:"isUnlisted"`      //IsUnlisted 是否为非公开作品
			IsMasked        bool      `json:"isMasked"`        //IsMasked 作品是否被屏蔽,即是否被隐藏或遮盖了一部分内容,如果这个字段为 true,则表示作品被屏蔽了,用户需要登录并且满足一定的条件才能查看这个作品,如果这个字段为 false,则表示作品没有被屏蔽,任何用户都可以查看
			AiType          int       `json:"aiType"`          //AiType 是否为AI作品
			SeriesID        string    `json:"seriesId"`        //SeriesID 作品系列ID
			SeriesTitle     string    `json:"seriesTitle"`     //SeriesTitle 作品系列标题
			ProfileImageURL string    `json:"profileImageUrl"` //ProfileImageURL 用户头像链接
		} `json:"userIllusts"` //UserIllusts 用户作品列表,只会返回几个含值项
		LikeData             bool   `json:"likeData"`      //LikeData 表示当前用户是否已经给该作品点赞
		Width                int    `json:"width"`         //Width 表示作品的宽度
		Height               int    `json:"height"`        //Height 表示作品的高度
		PageCount            int    `json:"pageCount"`     //PageCount 表示作品的页数
		BookmarkCount        int    `json:"bookmarkCount"` //BookmarkCount 表示作品被收藏的次数
		LikeCount            int    `json:"likeCount"`     //LikeCount 表示作品被点赞的次数
		CommentCount         int    `json:"commentCount"`  //CommentCount 表示作品被评论的次数
		ResponseCount        int    `json:"responseCount"` //ResponseCount 表示作品被回复的次数
		ViewCount            int    `json:"viewCount"`     //ViewCount 表示作品被浏览的次数
		BookStyle            string `json:"bookStyle"`     //BookStyle 表示作品的画集类型
		IsHowto              bool   `json:"isHowto"`       //IsHowto 表示作品是否为教程/指南类作品,如果为true,则表示该作品是一篇教程/指南,否则为 false
		IsOriginal           bool   `json:"isOriginal"`    //IsOriginal 表示这个作品是否为原创作品,如果为 true,则表示这个作品是作者原创的,如果为 false,则表示这个作品不是作者原创的,例如二次创作作品、合作作品等
		ImageResponseOutData []struct {
			Type     string `json:"type"`     //Type 作品类型[illust,manga,ugoira]
			WorkID   string `json:"workId"`   //WorkID 图片的ID
			Title    string `json:"title"`    //Title 图片的标题
			UserName string `json:"userName"` //UserName 图片作者的用户名
			ImageURL string `json:"imageUrl"` //ImageURL 图片的URL地址
		} `json:"imageResponseOutData"` //ImageResponseOutData 作品的是某作品的关联作品
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
				ID      string `json:"id"`      //ID 该收藏记录的唯一标识符,可以用于后续的修改或删除操作
				Private bool   `json:"private"` //Private 是否为私人收藏,如果为 true,则表示这个收藏仅对用户自己可见,如果为 false,则表示这个收藏为公开的
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
		} `json:"imageResponseData"` //ImageResponseData 作品的关联作品信息
		ImageResponseCount int         `json:"imageResponseCount"` //ImageResponseCount 关联作品的数量
		PollData           interface{} `json:"pollData"`           //PollData 投票信息
		SeriesNavData      struct {
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
		} `json:"seriesNavData"` //SeriesNavData 作品是否属于一个系列
		DescriptionBoothID   int         `json:"descriptionBoothId"`   //DescriptionBoothID BoothID的展位ID
		DescriptionYoutubeID int         `json:"descriptionYoutubeId"` //DescriptionYoutubeID Youtube的展位ID
		ComicPromotion       interface{} `json:"comicPromotion"`       //ComicPromotion 作品竞赛信息
		FanboxPromotion      struct {
			UserName        string `json:"userName"`        //UserName 作者名称
			UserImageURL    string `json:"userImageUrl"`    //UserImageURL 作者的头像url
			ContentURL      string `json:"contentUrl"`      //ContentURL Fanbox的主页URL
			Description     string `json:"description"`     //Description Fanbox的描述
			ImageURL        string `json:"imageUrl"`        //ImageURL Fanbox的封面图片URL
			ImageURLMobile  string `json:"imageUrlMobile"`  //ImageURLMobile Fanbox的移动版封面图片URL
			HasAdultContent bool   `json:"hasAdultContent"` //HasAdultContent Fanbox是否含有成人内容,如果含有则为true,否则为false
		} `json:"fanboxPromotion"` //FanboxPromotion Fanbox信息
		ContestBanners struct {
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
		} `json:"contestBanners"` //ContestBanners pixiv竞赛横幅相关信息
		IsBookmarkable bool `json:"isBookmarkable"` //IsBookmarkable 这个字段表示此作品是否允许用户加入收藏夹
		BookmarkData   struct {
			ID      string `json:"id"`      //ID 作品唯一书签ID
			Private bool   `json:"private"` //Private 是否为私人书签
		} `json:"bookmarkData"` //BookmarkData 这是有关书签数据的信息,包括是否已加入书签,收藏次数等
		ContestData struct {
			URL   string `json:"url"`   //URL 比赛链接地址(例:/contest/deformer_world2)
			Icon  string `json:"icon"`  //Icon 比赛内容图标
			Title string `json:"title"` //Title 比赛内容标题
		} `json:"contestData"` //ContestData 这是与比赛有关的数据,包括比赛ID,标题等
		ZoneConfig struct {
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
		} `json:"zoneConfig"` //ZoneConfig 表示绘画作品在不同展示区域（比如首页、排行榜、搜索结果页面等）显示时所使用的图片资源的URL配置
		ExtraData struct {
			Meta struct {
				Title              string `json:"title"`       //Title illust的标题
				Description        string `json:"description"` //Description illust的简介
				Canonical          string `json:"canonical"`   //Canonical illust的规范链接
				AlternateLanguages struct {
					Ja string `json:"ja"` //AlternateLanguages 日语版本
					En string `json:"en"` //AlternateLanguages 英语版本,响应会根据请求参数lang而变化(例如:lang=zn为中文)
				} `json:"alternateLanguages"` //AlternateLanguages 可用的其他语言版本
				DescriptionHeader string `json:"descriptionHeader"` //DescriptionHeader illust简介的标题
				Ogp               struct {
					Description string `json:"description"` //Description illust的Open Graph协议描述
					Image       string `json:"image"`       //Image illust的Open Graph协议图像链接
					Title       string `json:"title"`       //Title illust的Open Graph协议标题
					Type        string `json:"type"`        //Type illust的Open Graph协议类型
				} `json:"ogp"` //Ogp 包含用于Open Graph协议的元数据
				Twitter struct {
					Description string `json:"description"` //Description illust的Twitter描述
					Image       string `json:"image"`       //Image illust的Twitter图像链接
					Title       string `json:"title"`       //Title illust的Twitter标题
					Card        string `json:"card"`        //Card 用于定义illust在Twitter上的展示方式
				} `json:"twitter"` //Twitter 作者推特
			} `json:"meta"` //Meta 一个包含有关illust的元数据的结构体
		} `json:"extraData"` //ExtraData 额外数据
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
		IsUnlisted bool        `json:"isUnlisted"` //IsUnlisted 表示illust是否被设为不公开的（即未公开列表）
		Request    interface{} `json:"request"`    //Request 用于管理illust的请求
		CommentOff int         `json:"commentOff"` //CommentOff 表示illust的评论是否被关闭
		AiType     int         `json:"aiType"`     //AiType AI分级[0:非AI作品,1:色情图片识别模型,2:草稿模型,3:扫描动漫和插图的模型,4:扫描摄影和背景的模型,5:扫描普通的摄影和插图的模型,6:扫描儿童插图的模型]
	} `json:"body"` //Body pixiv返回的具体数据
}
