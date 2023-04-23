package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type UserRepo struct {
	DB *dao.DB
}

func NewUserRepo(db *dao.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) GetByID(id int64) (*entity.User, bool, error) {
	user := &entity.User{}
	exist, err := repo.DB.ID(id).Get(user)
	if err != nil {
		return nil, false, err
	}
	return user, exist, nil
}

func (repo *UserRepo) GetByName(name string) (*entity.User, bool, error) {
	user := &entity.User{}
	exist, err := repo.DB.Where("nickname = ?", name).Get(user)
	if err != nil {
		return nil, false, err
	}
	return user, exist, nil
}

func (repo *UserRepo) GetBySimilarName(name string, page, pageSize int) ([]*entity.User, int64, error) {
	var users []*entity.User
	total, err := repo.DB.Where("nickname LIKE ?", "%"+name+"%").
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&users)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (repo *UserRepo) Create(user *entity.User) error {
	_, err := repo.DB.InsertOne(user)
	return err
}

func (repo *UserRepo) Delete(uid int64) error {
	_, err := repo.DB.Where("uid = ?", uid).Delete(&entity.User{})
	return err
}
