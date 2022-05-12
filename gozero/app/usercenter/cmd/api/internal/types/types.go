// Code generated by goctl. DO NOT EDIT.
package types

type MemUser struct {
	Id              int64  `json:"id"`              //  id
	Username        string `json:"username"`        //用户名
	Password        string `json:"password"`        //密码
	PasswordSlat    string `json:"passwordSlat"`    //密码盐
	Email           string `json:"email"`           //邮件
	Mobile          string `json:"mobile"`          //手机
	Nickname        string `json:"nickname"`        //昵称
	Realname        string `json:"realname"`        //真实名
	CardId          string `json:"cardId"`          //身份证
	Sex             int    `json:"sex"`             //性别
	Birthday        *int64 `json:"birthday"`        //生日
	Avatar          string `json:"avatar"`          //头像
	MobileValidated bool   `json:"mobileValidated"` //验证手机
	EmailValidated  bool   `json:"emailValidated"`  //验证邮箱
	CardidValidated bool   `json:"cardidValidated"` //验证实名
	Info            string `json:"info"`            //备注
	RecommendCode   string `json:"recommendCode"`   //推荐码16位（自己的）
	Status          int    `json:"status"`          //状态
	StatusSafe      int    `json:"statusSafe"`      //安全状态
	RegIp           int    `json:"regIp"`           //注册ip
	LoginIp         int    `json:"loginIp"`         //登录ip
	LoginTotal      int    `json:"loginTotal"`      //登录次数
	LoginTime       *int64 `json:"loginTime"`       //最后登录时间
}

type RegisterReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type WXMiniAuthReq struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}

type WXMiniAuthResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo MemUser `json:"userInfo"`
}