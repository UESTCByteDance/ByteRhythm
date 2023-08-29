package dao

import (
	"ByteRhythm/model"
	"ByteRhythm/util"
	"context"

	"gorm.io/gorm"
)

type RelationDao struct {
	*gorm.DB
}

func NewRelationDao(ctx context.Context) *RelationDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &RelationDao{NewDBClient(ctx)}
}

func (r *RelationDao) FindUserById(uid int) (user *model.User, err error) {
	err = r.Model(&model.User{}).Where("id = ?", uid).Limit(1).Find(&user).Error
	if err != nil {
		return
	}
	return
}

func (r *RelationDao) FindAllFollow(uid int) (follows []*model.Follow, err error) {
	err = r.Model(&model.Follow{}).Where("followed_user_id = ?", uid).Find(&follows).Error
	if err != nil {
		return
	}
	return follows, nil
}

func (r *RelationDao) FindAllFollower(uid int) (follows []*model.Follow, err error) {
	err = r.Model(&model.Follow{}).Where("user_id = ?", uid).Find(&follows).Error
	if err != nil {
		return
	}
	return follows, nil
}

func (r *RelationDao) AddFollow(follow *model.Follow) (id int64, err error) {
	result := r.Model(&model.Follow{}).Where("user_id = ?", follow.UserID).Where("followed_user_id=?", follow.FollowedUserID).FirstOrCreate(&follow)
	if result.Error != nil {
		return
	}
	return result.RowsAffected, nil
}

func (r *RelationDao) CancelFollow(follow *model.Follow) (relation *model.Follow, err error) {
	err = r.Model(&model.Follow{}).Where("user_id = ?", follow.UserID).Where("followed_user_id=?", follow.FollowedUserID).Delete(&follow).Error
	if err != nil {
		return
	}
	return
}

func (r *RelationDao) GetFriendCount(userId int, followedUserId int) (count int64, err error) {
	err = r.Model(&model.Follow{}).Where("user_id = ?", userId).Where("followed_user_id=?", followedUserId).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (r *RelationDao) GetFollowCount(uid int) (count int64, err error) {
	err = r.Model(&model.Follow{}).Where("followed_user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (r *RelationDao) GetFollowerCount(uid int) (count int64, err error) {
	err = r.Model(&model.Follow{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (r *RelationDao) GetWorkCount(uid int) (count int64, err error) {
	err = r.Model(&model.Video{}).Where("author_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (r *RelationDao) GetFavoriteCount(uid int) (count int64, err error) {
	err = r.Model(&model.Favorite{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (r *RelationDao) GetTotalFavorited(uid int) (count int64, err error) {
	var videos []*model.Video
	err = r.Model(&model.Video{}).Where("author_id = ?", uid).Find(&videos).Error
	if err != nil {
		return
	}
	for _, video := range videos {
		var favoriteCount int64
		err = r.Model(&model.Favorite{}).Where("video_id = ?", video.ID).Count(&favoriteCount).Error
		if err != nil {
			return
		}
		count += favoriteCount
	}
	return
}

func (r *RelationDao) GetIsFollowed(uid int, token string) (isFollowed bool, err error) {

	baseId, err := util.GetUserIdFromToken(token)
	if err != nil {
		return
	}
	var follow model.Follow
	err = r.Model(&model.Follow{}).Where("user_id = ?", uid).Where("followed_user_id = ?", baseId).Limit(1).Find(&follow).Error
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
