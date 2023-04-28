package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
	"about-vaccine/internal/service/user"
)

type UserRepo struct {
	DB *dao.DB
}

func NewUserRepo(db *dao.DB) user.UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) GetByName(name string) (*entity.User, bool, error) {
	user := &entity.User{}
	exist, err := repo.DB.Where("nickname = ?", name).Get(user)
	if err != nil {
		return nil, false, err
	}
	return user, exist, nil
}

func (repo *UserRepo) GetListBySimilarName(name string, page, pageSize int) ([]*entity.User, int64, error) {
	var users []*entity.User
	total, err := repo.DB.Where("nickname LIKE ?", "%"+name+"%").Cols("id", "nickname").
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&users)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (repo *UserRepo) GetUserName(uid int64) (string, error) {
	var u entity.User
	_, err := repo.DB.Where("id = ?", uid).Cols("nickname").Get(&u)
	if err != nil {
		return "", err
	}
	return u.Nickname, nil
}

func (repo *UserRepo) Create(user *entity.User) error {
	_, err := repo.DB.InsertOne(user)
	return err
}

func (repo *UserRepo) Delete(uid int64) error {
	_, err := repo.DB.Where("id = ?", uid).Delete(&entity.User{})
	return err
}
