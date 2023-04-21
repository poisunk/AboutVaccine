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

func (repo *UserRepo) GetListByName(name string, page, pageSize int) ([]*entity.User, error) {
	var users []*entity.User
	err := repo.DB.Where("nickname = ?", name).Limit(pageSize, (page-1)*pageSize).Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepo) Create(user ...*entity.User) error {
	_, err := repo.DB.Insert(user)
	return err
}

func (repo *UserRepo) Delete(id int64) error {
	_, err := repo.DB.ID(id).Delete(&entity.User{})
	return err
}
