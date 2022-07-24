package models

//定义请求的参数的结构体
//ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}

//ParamLogin 登陆请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//投票数据
type ParamVoteData struct {
	//UserID
	PostID    string `json:"post_id,string" binding:"required"`                //帖子id
	Direction int8   `json:"direction,string" binging:"required,oneof=1 0 -1"` // 赞成票还是反对票
}
