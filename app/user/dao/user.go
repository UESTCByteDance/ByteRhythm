package dao

import (
	"ByteRhythm/model"
	model2 "ByteRhythm/model"
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

func (u *UserDao) FindUserByUserName(username string) (user *model2.User, err error) {
	//查看该用户是否存在
	err = u.Model(&model2.User{}).Where("username = ?", username).Limit(1).Find(&user).Error
	if err != nil {
		return
	}
	return
}
func (u *UserDao) CreateUser(user *model2.User) (id int64, err error) {
	err = u.Model(&model2.User{}).Create(&user).Error
	if err != nil {
		return
	}
	return int64(user.Id), nil
}

func (u *UserDao) FindUserById(id int64) (user *model2.User, err error) {
	err = u.Model(&model2.User{}).Where("id = ?", id).Limit(1).Find(&user).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetFollowCount(userId int64) (count int64, err error) {
	err = u.Model(&model.Follow{}).Where("followed_user_id = ?", userId).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetFollowerCount(userId int64) (count int64, err error) {
	err = u.Model(&model.Follow{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetWorkCount(userId int64) (count int64, err error) {
	err = u.Model(&model.Video{}).Where("author_id = ?", userId).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetFavoriteCount(userId int64) (count int64, err error) {
	err = u.Model(&model.Favorite{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (u *UserDao) GetTotalFavorited(userId int64) (count int64, err error) {
	var videos []*model.Video
	err = u.Model(&model.Video{}).Where("author_id = ?", userId).Find(&videos).Error
	if err != nil {
		return
	}
	for _, video := range videos {
		var favoriteCount int64
		err = u.Model(&model.Favorite{}).Where("video_id = ?", video.Id).Count(&favoriteCount).Error
		if err != nil {
			return
		}
		count += favoriteCount
	}
	return
}

func (u *UserDao) GetIsFollowed(userId int64, token string) (isFollowed bool, err error) {

	baseId, err := util.GetUserIdFromToken(token)
	if err != nil {
		return
	}
	var follow model.Follow
	err = u.Model(&model.Follow{}).Where("user_id = ?", baseId).Where("followed_user_id = ?", userId).Limit(1).Find(&follow).Error
	if err != nil {
		return
	}
	if follow.Id != 0 {
		isFollowed = true
	} else {
		isFollowed = false
	}
	return
}
