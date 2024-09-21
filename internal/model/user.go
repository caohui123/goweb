package model

type User struct {
	UserName string `json:"username" binding:"required,min=4,max=12" label:"用户名"`
	PassWord string `json:"password" binding:"required,min=6,max=20" label:"密码"`
}
