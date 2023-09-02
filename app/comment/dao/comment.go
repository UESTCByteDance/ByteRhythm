package dao

import (
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"

	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

func NewCommentDao(ctx context.Context) *CommentDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CommentDao{NewDBClient(ctx)}
}

func (c *CommentDao) CreateComment(comment *model.Comment) (err error) {
	err = c.Model(&model.Comment{}).Create(&comment).Error
	if err != nil {
		return
	}
	return nil
}

func (c *CommentDao) DeleteComment(comment *model.Comment) (err error) {
	err = c.Model(&model.Comment{}).Where(&comment).Delete(&comment).Error
	if err != nil {
		return
	}
	return nil
}

func (c *CommentDao) GetCommentListByVideoId(vid int64) (comments []*model.Comment, err error) {
	err = c.Model(&model.Comment{}).Where("video_id = ?", vid).Order("id desc").Find(&comments).Error
	if err != nil {
		return
	}
	return
}

func (c *CommentDao) GetUsernameByUid(uid int64) (username string, err error) {
	err = c.Model(&model.User{}).Where("id = ?", uid).Select("username").Find(&username).Error
	if err != nil {
		return
	}
	return
}

func (c *CommentDao) GetAvatarByUid(uid int64) (avatar string, err error) {
	err = c.Model(&model.User{}).Where("id = ?", uid).Select("avatar").Find(&avatar).Error
	if err != nil {
		return
	}
	return
}

func (c *CommentDao) GetFollowCount(uid int) (count int64, err error) {
	err = c.Model(&model.Follow{}).Where("followed_user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (c *CommentDao) GetFollowerCount(uid int) (count int64, err error) {
	err = c.Model(&model.Follow{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (c *CommentDao) GetWorkCount(uid int) (count int64, err error) {
	err = c.Model(&model.Video{}).Where("author_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (c *CommentDao) GetUserFavoriteCount(uid int) (count int64, err error) {
	err = c.Model(&model.Favorite{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (c *CommentDao) GetFavoriteCount(vid int) (count int64, err error) {
	err = c.Model(&model.Favorite{}).Where("video_id = ?", vid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (c *CommentDao) GetTotalFavorited(uid int) (count int64, err error) {
	var videos []*model.Video
	err = c.Model(&model.Video{}).Where("author_id = ?", uid).Find(&videos).Error
	if err != nil {
		return
	}
	for _, video := range videos {
		var favoriteCount int64
		err = c.Model(&model.Favorite{}).Where("video_id = ?", video.ID).Count(&favoriteCount).Error
		if err != nil {
			return
		}
		count += favoriteCount
	}
	return
}

func (c *CommentDao) GetIsFollowed(uid int, token string) (isFollowed bool, err error) {

	baseID, err := util.GetUserIdFromToken(token)
	if err != nil {
		return
	}
	var follow model.Follow
	err = c.Model(&model.Follow{}).Where("user_id = ?", baseID).Where("followed_user_id = ?", uid).Limit(1).Find(&follow).Error
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
