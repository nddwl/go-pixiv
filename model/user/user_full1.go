package user

//UserFull1 https://www.pixiv.net/ajax/user/<user_id>?full=1&lang=<lang>&version=<version>
type UserFull1 struct {
	Error   bool   `json:"error"`   // Error 表示是否有错误,如果没有错误则为false,有错误则为true
	Message string `json:"message"` //Message 错误信息
	Body    struct {
		UserID     string `json:"userId"`     //UserID 用户ID
		Name       string `json:"name"`       //Name 用户名
		Image      string `json:"image"`      //Image 用户头像URL
		ImageBig   string `json:"imageBig"`   //ImageBig 用户头像URL（高清）
		Premium    bool   `json:"premium"`    //Premium 表示用户是否为高级会员,如果是则为true,否则为false
		IsFollowed bool   `json:"isFollowed"` //IsFollowed 表示当前登录的用户是否已经关注了这个用户,如果是则为true,否则为false
		IsMypixiv  bool   `json:"isMypixiv"`  //IsMypixiv 表示这个用户是否是当前登录用户的mypixiv,如果是则为true,否则为false
		IsBlocking bool   `json:"isBlocking"` //IsBlocking 表示当前登录用户是否被屏蔽,如果是则为true,否则为false
		Background struct {
			Repeat    bool   `json:"repeat"`    //Repeat 是否平铺背景,如果不是则为nil
			Color     string `json:"color"`     //Color 背景颜色,如果没有背景颜色则为nil
			URL       string `json:"url"`       //URL 背景图片URL
			IsPrivate bool   `json:"isPrivate"` //IsPrivate 背景是否为私人的,如果是则为true,否则为false
		} `json:"background"` //Background 包含用户的背景信息
		SketchLiveID  string `json:"sketchLiveId"`  //SketchLiveID 用于Live Sketch功能的ID,如果没有则为nil
		Partial       int    `json:"partial"`       //Partial 如果用户的资料是不完整的,这里会显示不完整的信息数量
		AcceptRequest bool   `json:"acceptRequest"` //AcceptRequest 表示用户是否接受私信请求,如果接受则为true,否则为false
		SketchLives   []struct {
			ID            string `json:"id"`            //ID 模型的唯一标识符
			Name          string `json:"name"`          //Name 模型的名称
			URL           string `json:"url"`           //URL 模型的链接
			ThumbnailURL  string `json:"thumbnailUrl"`  //ThumbnailURL 模型的缩略图链接
			AudienceCount int    `json:"audienceCount"` //AudienceCount 观众数量
			IsR18         bool   `json:"isR18"`         //IsR18 表示模型是否为 R18
			StreamerIds   []int  `json:"streamerIds"`   //StreamerIds 流媒体 ID 列表
		} `json:"sketchLives"` //SketchLives 用于Live Sketch功能,这里会显示用户已经创建的sketch lives列表,如果没有则为nil
		Following    int    `json:"following"`    //Following 用户关注的人数
		MypixivCount int    `json:"mypixivCount"` //MypixivCount mypixiv内的作品数量
		FollowedBack bool   `json:"followedBack"` //FollowedBack 表示当前登录用户是否被这个用户关注了,如果是则为true,否则为false
		Comment      string `json:"comment"`      //Comment 用户自我介绍
		CommentHTML  string `json:"commentHtml"`  //CommentHTML 格式化后的自我介绍HTML代码
		Webpage      string `json:"webpage"`      //Webpage 用户的个人主页URL
		Social       struct {
			Twitter struct {
				URL string `json:"url"` //URL Twitter链接
			} `json:"twitter"` //Twitter 用户的Twitter账号信息
			Pawoo struct {
				URL string `json:"url"` //URL Pawoo链接
			} `json:"pawoo"` //Pawoo 用户的Pawoo账号信息
		} `json:"social"` //Social 社交账号链接
		CanSendMessage bool     `json:"canSendMessage"` //CanSendMessage 表示当前登录用户是否可以给这个用户发送私信,如果可以则为true,否则为false
		Region         struct { // 地域信息
			Name         string `json:"name"`         //Name 地域名称
			Region       string `json:"region"`       //Region 地区代码
			Prefecture   string `json:"prefecture"`   //Prefecture 省份信息,如果没有则为nil
			PrivacyLevel string `json:"privacyLevel"` //PrivacyLevel 地区信息的隐私等级
		} `json:"region"` //Region 用户所在的地区信息
		Age struct {
			Name         string `json:"name"`         //Name 年龄段名称
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
			UserWorkspaceTool   string `json:"userWorkspaceTool"`   //UserWorkspaceTool 用户的工作区工具
			UserWorkspaceTablet string `json:"userWorkspaceTablet"` //UserWorkspaceTablet 用户的工作区平板电脑
		} `json:"workspace"` //Workspace 包含用户的工作区信息
		Official bool `json:"official"` //Official 表示用户是否为官方账号
		Group    []struct {
			ID      string `json:"id"`      //ID 组的唯一标识符
			Title   string `json:"title"`   //Title 组的标题
			IconURL string `json:"iconUrl"` //IconURL 组的图标 URL
		} `json:"group"` //Group 表示用户所属的组信息
	} `json:"body"` //Body pixiv返回的具体数据
}
