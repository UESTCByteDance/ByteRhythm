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

func (f *FavoriteDao) CreateFavorite(favorite *model.Favorite) (err error) {
	err = f.Model(&model.Favorite{}).Create(&favorite).Error
	if err != nil {
		return
	}
	return nil
}

func (f *FavoriteDao) DeleteFavorite(favorite *model.Favorite) (err error) {
	err = f.Model(&model.Favorite{}).Where(&favorite).Delete(&favorite).Error
	if err != nil {
		return
	}
	return nil
}

func (f *FavoriteDao) GetFavoriteListByUserId(uid int64) (favorites []*model.Favorite, err error) {
	err = f.Model(&model.Favorite{}).Where("user_id = ?", uid).Find(&favorites).Error
	if err != nil {
		return
	}
	return
}

func (f *FavoriteDao) GetIsFavoriteByUserIdAndVid(uid int64, vid int64) (isFavorite bool, err error) {
	var count int64
	err = f.Model(&model.Favorite{}).Where("user_id = ?", uid).Where("video_id = ?", vid).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (f *FavoriteDao) GetFollowCount(uid int) (count int64, err error) {
	err = f.Model(&model.Follow{}).Where("followed_user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (f *FavoriteDao) GetFollowerCount(uid int) (count int64, err error) {
	err = f.Model(&model.Follow{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (f *FavoriteDao) GetWorkCount(uid int) (count int64, err error) {
	err = f.Model(&model.Video{}).Where("author_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (f *FavoriteDao) GetUserFavoriteCount(uid int) (count int64, err error) {
	err = f.Model(&model.Favorite{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (f *FavoriteDao) GetFavoriteCount(vid int) (count int64, err error) {
	err = f.Model(&model.Favorite{}).Where("video_id = ?", vid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (f *FavoriteDao) GetTotalFavorited(uid int) (count int64, err error) {
	var videos []*model.Video
	err = f.Model(&model.Video{}).Where("author_id = ?", uid).Find(&videos).Error
	if err != nil {
		return
	}
	for _, video := range videos {
		var favoriteCount int64
		err = f.Model(&model.Favorite{}).Where("video_id = ?", video.ID).Count(&favoriteCount).Error
		if err != nil {
			return
		}
		count += favoriteCount
	}
	return
}

func (f *FavoriteDao) GetIsFollowed(uid int, token string) (isFollowed bool, err error) {

	baseID, err := util.GetUserIdFromToken(token)
	if err != nil {
		return
	}
	var follow model.Follow
	err = f.Model(&model.Follow{}).Where("user_id = ?", baseID).Where("followed_user_id = ?", uid).Limit(1).Find(&follow).Error
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

func (f *FavoriteDao) GetPlayUrlByVid(vid int) (playUrl string, err error) {
	err = f.Model(&model.Video{}).Where("id = ?", vid).Select("play_url").First(&playUrl).Error
	if err != nil {
		return
	}
	return
}

func (f *FavoriteDao) GetCoverUrlByVid(vid int) (playUrl string, err error) {
	err = f.Model(&model.Video{}).Where("id = ?", vid).Select("cover_url").First(&playUrl).Error
	if err != nil {
		return
	}
	return
}
