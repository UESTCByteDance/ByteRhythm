package dao

import (
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"
	"time"

	"gorm.io/gorm"
)

type VideoDao struct {
	*gorm.DB
}

func NewVideoDao(ctx context.Context) *VideoDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &VideoDao{NewDBClient(ctx)}
}

func (v *VideoDao) GetVideoListByLatestTime(latestTime time.Time, vidExistArray []int, limit int) (videos []*model.Video, err error) {
	if limit != 30 {
		err = v.Model(&model.Video{}).Where("created_at <= ?", latestTime).Not("id", vidExistArray).Order("created_at desc").Limit(limit).Find(&videos).Error
		if err != nil {
			return
		}
	}
	err = v.Model(&model.Video{}).Where("created_at <= ?", latestTime).Order("created_at desc").Limit(limit).Find(&videos).Error
	if err != nil {
		return
	}
	return
}

func (v *VideoDao) GetFavoriteCount(vid int) (count int64, err error) {
	err = v.Model(&model.Favorite{}).Where("video_id = ?", vid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *VideoDao) GetCommentCount(vid int) (count int64, err error) {
	err = v.Model(&model.Comment{}).Where("video_id = ?", vid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *VideoDao) GetIsFavorite(vid int, token string) (isFavorite bool, err error) {
	baseID, err := util.GetUserIdFromToken(token)
	if err != nil {
		return
	}
	var favorite model.Favorite
	err = v.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", baseID, vid).Limit(1).Find(&favorite).Error
	if err != nil {
		return
	}
	if favorite.ID != 0 {
		isFavorite = true
	} else {
		isFavorite = false
	}
	return
}

func (v *VideoDao) GetFollowCount(uid int) (count int64, err error) {
	err = v.Model(&model.Follow{}).Where("followed_user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *VideoDao) GetFollowerCount(uid int) (count int64, err error) {
	err = v.Model(&model.Follow{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *VideoDao) GetWorkCount(uid int) (count int64, err error) {
	err = v.Model(&model.Video{}).Where("author_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *VideoDao) GetUserFavoriteCount(uid int) (count int64, err error) {
	err = v.Model(&model.Favorite{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (v *VideoDao) GetTotalFavorited(uid int) (count int64, err error) {
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

func (v *VideoDao) GetIsFollowed(uid int, token string) (isFollowed bool, err error) {

	baseID, err := util.GetUserIdFromToken(token)
	if err != nil {
		return
	}
	var follow model.Follow
	err = v.Model(&model.Follow{}).Where("user_id = ?", uid).Where("followed_user_id = ?", baseID).Limit(1).Find(&follow).Error
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

func (v *VideoDao) GetVideoListByUserId(uid int) (videos []*model.Video, err error) {
	err = v.Model(&model.Video{}).Where("author_id = ?", uid).Order("created_at desc").Limit(30).Find(&videos).Error
	if err != nil {
		return
	}
	return
}

func (v *VideoDao) CreateVideo(video *model.Video) (err error) {
	err = v.Model(&model.Video{}).Create(&video).Error
	if err != nil {
		return
	}
	return
}

func (v *VideoDao) FindUser(video *model.Video) (user *model.User, err error) {
	//gorm通过外键查询
	err = v.Model(&video).Association("Author").Find(&user)
	if err != nil {
		return
	}
	return
}
