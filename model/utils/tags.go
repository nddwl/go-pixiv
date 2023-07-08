package utils

//Tags `json:"tags"`
type Tags struct {
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
}
