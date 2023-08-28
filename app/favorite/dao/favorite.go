package dao

import (
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"

	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &FavoriteDao{NewDBClient(ctx)}
}

func (c *FavoriteDao) CreateFavorite(favorite *model.Favorite) (err error) {
	err = c.Model(&model.Favorite{}).Create(&favorite).Error
	if err != nil {
		return
	}
	return nil
}

func (c *FavoriteDao) DeleteFavorite(favorite *model.Favorite) (err error) {
	err = c.Model(&model.Favorite{}).Where(&favorite).Delete(&favorite).Error
	if err != nil {
		return
	}
	return nil
}

func (c *FavoriteDao) GetFavoriteListByUserId(uid int64) (favorites []*model.Favorite, err error) {
	err = c.Model(&model.Favorite{}).Where("user_id = ?", uid).Find(&favorites).Error
	if err != nil {
		return
	}
	return
}

func (v *FavoriteDao) GetFollowCount(uid int) (count int64, err error) {
	err = v.Model(&model.Follow{}).Where("followed_user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *FavoriteDao) GetFollowerCount(uid int) (count int64, err error) {
	err = v.Model(&model.Follow{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *FavoriteDao) GetWorkCount(uid int) (count int64, err error) {
	err = v.Model(&model.Video{}).Where("author_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *FavoriteDao) GetUserFavoriteCount(uid int) (count int64, err error) {
	err = v.Model(&model.Favorite{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *FavoriteDao) GetFavoriteCount(vid int) (count int64, err error) {
	err = v.Model(&model.Favorite{}).Where("video_id = ?", vid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *FavoriteDao) GetTotalFavorited(uid int) (count int64, err error) {
	var videos []*model.Video
	err = v.Model(&model.Video{}).Where("author_id = ?", uid).Find(&videos).Error
	if err != nil {
		return
	}
	for _, video := range videos {
		var favoriteCount int64
		err = v.Model(&model.Favorite{}).Where("video_id = ?", video.ID).Count(&favoriteCount).Error
		if err != nil {
			return
		}
		count += favoriteCount
	}
	return
}

func (v *FavoriteDao) GetIsFollowed(uid int, token string) (isFollowed bool, err error) {

	baseID, err := util.GetUserIdFromToken(token)
	if err != nil {
		return
	}
	var follow model.Follow
	err = v.Model(&model.Follow{}).Where("user_id = ?", baseID).Where("followed_user_id = ?", uid).Limit(1).Find(&follow).Error
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

func (v *FavoriteDao) GetPlayUrlByVid(vid int) (playUrl string, err error) {
	err = v.Model(&model.Video{}).Where("id = ?", vid).Select("play_url").First(&playUrl).Error
	if err != nil {
		return
	}
	return
}

func (v *FavoriteDao) GetCoverUrlByVid(vid int) (playUrl string, err error) {
	err = v.Model(&model.Video{}).Where("id = ?", vid).Select("cover_url").First(&playUrl).Error
	if err != nil {
		return
	}
	return
}
