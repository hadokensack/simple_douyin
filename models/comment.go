package models

import "time"

type Comment struct {
	Id         int64     `json:"id"`            // 评论id
	User       UserInfo  `json:"user" gorm:"-"` // 评论用户信息
	CreatedAt  time.Time `json:"-"`             //评论创建时间
	UserInfoId int64     `json:"-"`             //用于一对多关系的id
	VideoId    int64     `json:"-"`             //一对多，视频对评论
}
