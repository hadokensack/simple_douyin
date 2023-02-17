package service

import (
	"errors"
	"simpel_douyin/models"
)

type List struct {
	Videos []*models.Video `json:"video_list,omitempty"`
}

func QueryVideoListByUserId(userId int64) (*List, error) {
	return NewQueryVideoListByUserIdFlow(userId).Do()
}

func NewQueryVideoListByUserIdFlow(userId int64) *QueryVideoListByUserIdFlow {
	return &QueryVideoListByUserIdFlow{userId: userId}
}

type QueryVideoListByUserIdFlow struct {
	userId int64
	videos []*models.Video

	videoList *List
}

func (q *QueryVideoListByUserIdFlow) Do() (*List, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.videoList, nil
}

func (q *QueryVideoListByUserIdFlow) checkNum() error {
	//检查userId是否存在
	if !models.NewUserInfoDAO().IsUserExistById(q.userId) {
		return errors.New("用户不存在")
	}

	return nil
}

// 填充数据到数据库
func (q *QueryVideoListByUserIdFlow) packData() error {
	err := models.NewVideoDAO().QueryVideoListByUserId(q.userId, &q.videos)
	if err != nil {
		return err
	}

	return nil
}
