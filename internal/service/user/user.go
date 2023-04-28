package user

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
)

type UserRepo interface {
	GetListBySimilarName(keyword string, page, pageSize int) ([]*entity.User, int64, error)
	GetUserName(uid int64) (string, error)
	GetByName(name string) (*entity.User, bool, error)
	Create(user *entity.User) error
	Delete(uid int64) error
}

type UserCommon struct {
	UserRepo UserRepo
}

func NewUserCommon(userRepo UserRepo) *UserCommon {
	return &UserCommon{
		UserRepo: userRepo,
	}
}

func (u *UserCommon) GetListBySimilarName(keyword string, page, pageSize int) ([]*schema.UserInfo, int64, error) {
	users, total, err := u.UserRepo.GetListBySimilarName(keyword, page, pageSize)
	if err != nil {
		return nil, total, err
	}
	userList := make([]*schema.UserInfo, 0)
	for _, v := range users {
		user := u.FormatInfo(v)
		userList = append(userList, user)
	}
	return userList, total, nil
}

func (u *UserCommon) GetUserNameByUid(uid int64) (string, error) {
	userName, err := u.UserRepo.GetUserName(uid)
	if err != nil {
		return "", err
	}
	return userName, nil
}

func (u *UserCommon) GetByName(name string) (*entity.User, bool, error) {
	return u.UserRepo.GetByName(name)
}

func (u *UserCommon) CheckName(name string) (bool, error) {
	_, exist, err := u.UserRepo.GetByName(name)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (u *UserCommon) Create(user *schema.UserAdd) (int64, error) {
	uEntity := &entity.User{
		Nickname: user.Nickname,
		Password: user.Password,
	}
	err := u.UserRepo.Create(uEntity)
	return uEntity.Id, err
}

func (u *UserCommon) Delete(uid int64) error {
	return u.UserRepo.Delete(uid)
}

func (u *UserCommon) FormatInfo(user *entity.User) *schema.UserInfo {
	return &schema.UserInfo{
		Uid:      user.Id,
		Nickname: user.Nickname,
	}
}
