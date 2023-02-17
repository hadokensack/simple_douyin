package models

import (
	"errors"
	"log"
	"sync"
)

var (
	ErrIvdPtr        = errors.New("空指针错误")
	ErrEmptyUserList = errors.New("用户列表为空")
)

type UserInfo struct {
	Id            int64      `json:"id" gorm:"id,omitempty"`
	Name          string     `json:"name" gorm:"name,omitempty"`
	FollowCount   int64      `json:"follow_count" gorm:"follow_count,omitempty"`
	FollowerCount int64      `json:"follower_count" gorm:"follower_count,omitempty"`
	IsFollow      bool       `json:"is_follow" gorm:"is_follow,omitempty"`
	User          *UserLogin `json:"-"` //用户与账号密码之间的一对一
	Videos        []*Video   `json:"-"` //用户与投稿视频的一对多
	Comments      []*Comment `json:"-"` //用户与评论的一对多
}

type UserInfoDAO struct {
}

var (
	userInfoDAO  *UserInfoDAO
	userInfoOnce sync.Once
)

func NewUserInfoDAO() *UserInfoDAO {
	userInfoOnce.Do(func() {
		userInfoDAO = new(UserInfoDAO)
	})
	return userInfoDAO
}

func (u *UserInfoDAO) QueryUserInfoById(userId int64, userinfo *UserInfo) error {
	if userinfo == nil {
		return ErrIvdPtr
	}
	//DB.Where("id=?",userId).First(userinfo)
	DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
	//id为零值，说明sql执行失败
	if userinfo.Id == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

// 增加用户信息
func (u *UserInfoDAO) AddUserInfo(userinfo *UserInfo) error {
	if userinfo == nil {
		return ErrIvdPtr
	}
	return DB.Create(userinfo).Error
}

// 用户是否存在
func (u *UserInfoDAO) IsUserExistById(id int64) bool {
	var userinfo UserInfo
	if err := DB.Where("id=?", id).Select("id").First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.Id == 0 {
		return false
	}
	return true
}
