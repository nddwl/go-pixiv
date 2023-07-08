package user

import (
	"go-pixiv/model/utils"
)

//ProfileTop https://www.pixiv.net/ajax/user/<user_id>/profile/top?lang=<lang>&version=<version>
type ProfileTop struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		Illusts     []utils.Work `json:"illusts"`
		Manga       []utils.Work `json:"manga"`
		Novels      []utils.Work `json:"novels"`
		RequestWork struct {
			Artworks []interface{} `json:"artworks"`
			Novels   []interface{} `json:"novels"`
		} `json:"requestPostWorks"`
		RequestPlans []interface{} `json:"requestPlans"`
		ZoneConfig   struct {
			Header struct {
				URL string `json:"url"` //URL 链接
			} `json:"header"` //Header 页眉展示区域的图片URL
			Footer struct {
				URL string `json:"url"` //URL 链接
			} `json:"footer"` //Footer 页脚展示区域的图片URL
			Logo struct {
				URL string `json:"url"` //URL 链接
			} `json:"logo"` //Logo Pixiv标志展示区域的图片URL
			Five00X500 struct {
				URL string `json:"url"` //URL 链接
			} `json:"500x500"` //Five00X500 500 x 500像素展示区域的图片URL
		} `json:"zoneConfig"`
		ExtraData struct {
			Meta struct {
				Title       string `json:"title"`       //Title illust的标题
				Description string `json:"description"` //Description illust的简介
				Canonical   string `json:"canonical"`   //Canonical illust的规范链接
				Ogp         struct {
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
				AlternateLanguages struct {
					Ja string `json:"ja"` //AlternateLanguages 日语版本
					En string `json:"en"` //AlternateLanguages 英语版本,响应会根据请求参数lang而变化(例如:lang=zn为中文)
				} `json:"alternateLanguages"` //AlternateLanguages 可用的其他语言版本
				DescriptionHeader string `json:"descriptionHeader"` //DescriptionHeader illust简介的标题
			} `json:"meta"` //Meta 一个包含有关illust的元数据的结构体
		} `json:"extraData"`
	} `json:"body"`
}
