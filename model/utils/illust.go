package utils

import (
	"bytes"
	"encoding/json"
	"time"
)

type Work struct {
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
		WorkTitle struct {
			En string //En 根据用户的语言或lang来返回结果
			Ja string //Ja 默认
		} `json:"workTitle"` //WorkTitle 标题翻译
		WorkCaption struct {
			En string //En 根据用户的语言或lang来返回结果
			Ja string //Ja 默认
		} `json:"workCaption"` //WorkCaption 描述翻译
	} `json:"titleCaptionTranslation"` //作品标题和作品描述的翻译信息
	CreateDate      time.Time `json:"createDate"`      //CreateDate 上传时间
	UpdateDate      time.Time `json:"updateDate"`      //UpdateDate 更新时间
	IsUnlisted      bool      `json:"isUnlisted"`      //IsUnlisted 是否为非公开作品
	IsMasked        bool      `json:"isMasked"`        //IsMasked 作品是否被屏蔽,即是否被隐藏或遮盖了一部分内容,如果这个字段为 true,则表示作品被屏蔽了,用户需要登录并且满足一定的条件才能查看这个作品,如果这个字段为 false,则表示作品没有被屏蔽,任何用户都可以查看
	AiType          int       `json:"aiType"`          //AiType 是否为AI作品
	ProfileImageURL string    `json:"profileImageUrl"` //ProfileImageURL 用户头像链接
}

type Illust struct {
	Work []Work
}

func (t *Illust) UnmarshalJSON(data []byte) error {
	if t == nil {
		*t = Illust{}
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.Token()
	for decoder.More() {
		var work Work
		decoder.Token()
		decoder.Decode(&work)
		t.Work = append(t.Work, work)
	}
	return nil
}

type IllustId struct {
	Work []Work
	Id   []string
}

func (t *IllustId) UnmarshalJSON(data []byte) error {
	if t == nil {
		*t = IllustId{}
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.Token()
	for decoder.More() {
		var work Work
		key, _ := decoder.Token()
		decoder.Decode(&work)
		if work.ID == "" {
			t.Id = append(t.Id, key.(string))
			for decoder.More() {
				key, _ = decoder.Token()
				decoder.Token()
				t.Id = append(t.Id, key.(string))
			}
		} else {
			t.Work = append(t.Work, work)
		}
	}
	return nil
}
