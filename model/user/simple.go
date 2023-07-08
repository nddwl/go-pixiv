package user

//Simple https://www.pixiv.net/ajax/user/<user_id>?full=0&lang=<lang>&version=<version>
type Simple struct {
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
	} `json:"body"`
}
