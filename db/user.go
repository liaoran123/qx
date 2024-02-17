package db

//对应user表
type User struct {
	Tid       int `json:"tid"`       //用户id
	Psw       int `json:"-"`         //密码 该字段忽略
	Cdtime    int `json:"cdtime"`    //注册时间
	Lastdtime int `json:"lastdtime"` //最后登录时间
}

func NewUser() *User {
	return &User{}
}
func GetUser(tid int) *User {
	return &User{}
}
