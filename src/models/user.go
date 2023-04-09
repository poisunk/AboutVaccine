package models

import "MyWeb/dao"

type User struct {
	UId      int    `gorm:"int(8);column:uid;primary_key" json:"uid"`
	Nickname string `gorm:"varchar(255);column:nickname" json:"nickname"`
	Password string `gorm:"char(64);column:password" json:"password"`
}

func GetUserList(page int, pageSize int) (userList []*User, err error) {
	if err = dao.DB.Offset((page - 1) * pageSize).Limit(pageSize).
		Select("uid, nickname").Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func GetUserById(id int64) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Where("uid = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func GetUserByName(name string) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Where("nickname = ?", name).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func CreateUser(user *User) (err error) {
	if err = dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserById(id int64) (err error) {
	if err = dao.DB.Delete(User{}, "uid = ?", id).Error; err != nil {
		return err
	}
	return
}

func CountUser() (count int64, err error) {
	if err = dao.DB.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count - 1, nil
}
