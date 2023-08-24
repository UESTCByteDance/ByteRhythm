package dao

import (
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (u *UserDao) FindUserByUserName(username string) (user *model.User, err error) {
	//查看该用户是否存在
	err = u.Model(&model.User{}).Where("username = ?", username).Limit(1).Find(&user).Error
	if err != nil {
		return
	}
	return
}
func (u *UserDao) CreateUser(user *model.User) (id int64, err error) {
	err = u.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return
	}
	return int64(user.ID), nil
}

func (u *UserDao) FindUserById(uid int) (user *model.User, err error) {
	err = u.Model(&model.User{}).Where("id = ?", uid).Limit(1).Find(&user).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetFollowCount(uid int) (count int64, err error) {
	err = u.Model(&model.Follow{}).Where("followed_user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetFollowerCount(uid int) (count int64, err error) {
	err = u.Model(&model.Follow{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetWorkCount(uid int) (count int64, err error) {
	err = u.Model(&model.Video{}).Where("author_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetFavoriteCount(uid int) (count int64, err error) {
	err = u.Model(&model.Favorite{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetTotalFavorited(uid int) (count int64, err error) {
	var videos []*model.Video
	err = u.Model(&model.Video{}).Where("author_id = ?", uid).Find(&videos).Error
	if err != nil {
		return
	}
	for _, video := range videos {
		var favoriteCount int64
		err = u.Model(&model.Favorite{}).Where("video_id = ?", video.ID).Count(&favoriteCount).Error
		if err != nil {
			return
		}
		count += favoriteCount
	}
	return
}

func (u *UserDao) GetIsFollowed(uid int, token string) (isFollowed bool, err error) {

	baseID, err := util.GetUserIdFromToken(token)
	if err != nil {
		return
	}
	var follow model.Follow
	err = u.Model(&model.Follow{}).Where("user_id = ?", uid).Where("followed_user_id = ?", baseID).Limit(1).Find(&follow).Error
	if err != nil {
		return
	}
	if follow.ID != 0 {
		isFollowed = true
	} else {
		isFollowed = false
	}
	return
}
