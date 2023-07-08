package user

//Detailed https://www.pixiv.net/ajax/user/<user_id>?full=1&lang=<lang>&version=<version>
type Detailed struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		UserID     string `json:"userId"`     //UserID 用户ID
		Name       string `json:"name"`       //Name 用户名
		Image      string `json:"image"`      //Image 用户头像
		ImageBig   string `json:"imageBig"`   //ImageBig 用户大头像
		Premium    bool   `json:"premium"`    //Premium 是否为大会员
		IsFollowed bool   `json:"isFollowed"` //IsFollowed 是否关注
		IsMypixiv  bool   `json:"isMypixiv"`  //IsMypixiv 是否为好友
		IsBlocking bool   `json:"isBlocking"` //IsBlocking 是否被屏蔽
		Background struct {
			Repeat    bool   `json:"repeat"`    //Repeat 是否平铺背景
			Color     string `json:"color"`     //Color 背景颜色
			URL       string `json:"url"`       //URL 背景图片
			IsPrivate bool   `json:"isPrivate"` //IsPrivate 背景是否为私人
		} `json:"background"` //Background 背景信息
		SketchLiveID  string `json:"sketchLiveId"`  //SketchLiveID LiveSketchID
		Partial       int    `json:"partial"`       //Partial 不完整信息数量
		AcceptRequest bool   `json:"acceptRequest"` //AcceptRequest 是否接受请求
		SketchLives   []struct {
			ID            string `json:"id"`            //ID 模型ID
			Name          string `json:"name"`          //Name 模型名称
			URL           string `json:"url"`           //URL 模型链接
			ThumbnailURL  string `json:"thumbnailUrl"`  //ThumbnailURL 模型缩略图
			AudienceCount int    `json:"audienceCount"` //AudienceCount 观众数量
			IsR18         bool   `json:"isR18"`         //IsR18 表示模型是否为R18
			StreamerIds   []int  `json:"streamerIds"`   //StreamerIds 流媒体ID列表
		} `json:"sketchLives"` //SketchLives LiveSketch信息
		Following    int    `json:"following"`    //Following 关注人数
		MypixivCount int    `json:"mypixivCount"` //MypixivCount 好友数量
		FollowedBack bool   `json:"followedBack"` //FollowedBack 是否被关注
		Comment      string `json:"comment"`      //Comment 自我介绍
		CommentHTML  string `json:"commentHtml"`  //CommentHTML 自我介绍HTML代码
		Webpage      string `json:"webpage"`      //Webpage 个人主页
		Social       struct {
			Twitter struct {
				URL string `json:"url"` //URL Twitter链接
			} `json:"twitter"` //Twitter Twitter账号信息
			Pawoo struct {
				URL string `json:"url"` //URL Pawoo链接
			} `json:"pawoo"` //Pawoo Pawoo账号信息
		} `json:"social"` //Social 社交账号链接
		CanSendMessage bool     `json:"canSendMessage"` //CanSendMessage 是否可以私信
		Region         struct { // 地域信息
			Name         string `json:"name"`         //Name 名称
			Region       string `json:"region"`       //Region 代码
			Prefecture   string `json:"prefecture"`   //Prefecture 省份信息
			PrivacyLevel string `json:"privacyLevel"` //PrivacyLevel 隐私等级
		} `json:"region"` //Region 地区信息
		Age struct {
			Name         string `json:"name"`         //Name 年龄段
			PrivacyLevel string `json:"privacyLevel"` //PrivacyLevel 隐私等级
		} `json:"age"` //Age 年龄信息
		BirthDay struct {
			Name         string `json:"name"`         //Name 年龄名称
			PrivacyLevel string `json:"privacyLevel"` //PrivacyLevel 隐私等级
		} `json:"birthDay"` //BirthDay 用户的生日信息
		Gender struct {
			Name         string `json:"name"`         //Name 性别名称
			PrivacyLevel string `json:"privacyLevel"` //PrivacyLevel 隐私等级
		} `json:"gender"` //Gender 用户的性别信息
		Job struct {
			Name         string `json:"name"`         //Name 工作名称
			PrivacyLevel string `json:"privacyLevel"` //PrivacyLevel 隐私等级
		} `json:"job"` //Job 工作
		Workspace struct {
			UserWorkspaceTool   string `json:"userWorkspaceTool"`   //UserWorkspaceTool 工作软件
			UserWorkspaceTablet string `json:"userWorkspaceTablet"` //UserWorkspaceTablet 工作工具
		} `json:"workspace"` //Workspace 工作间
		Official bool `json:"official"` //Official 是否为官方账号
		Group    []struct {
			ID      string `json:"id"`      //ID 组的唯一标识符
			Title   string `json:"title"`   //Title 组的标题
			IconURL string `json:"iconUrl"` //IconURL 组的图标 URL
		} `json:"group"` //Group 所属组
	} `json:"body"` //Body pixiv返回的具体数据
}
