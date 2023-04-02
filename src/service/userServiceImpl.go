package service

import (
	"MyWeb/middleware/jwt"
	"MyWeb/models"
	"MyWeb/utile"
	"errors"
	"log"
	"strconv"
)

type UserServiceImpl struct {
}

func InitUserService() *UserServiceImpl {
	service := &UserServiceImpl{}
	return service
}

func (service *UserServiceImpl) Login(user User) (token string, err error) {
	// 检查用户名与密码
	if len(user.Nickname) == 0 || len(user.Password) == 0 {
		return "", errors.New("用户名或密码不能为空！")
	}
	// 获取用户
	u := new(models.User)
	if u, err = models.GetUserByName(user.Nickname); err != nil {
		log.Println(err.Error())
		return "", errors.New("找不到用户")
	}
	// 检查密码
	if encodePassword(user.Password) != u.Password {
		return "", errors.New("密码不正确！")
	}
	// 生成token
	if token, err = jwt.GenerateToken(int64(u.UID), u.Nickname); err != nil {
		log.Println(err.Error())
		return "", errors.New("token生成失败！")
	}
	return token, nil
}

func (service *UserServiceImpl) Register(user User) (token string, err error) {
	// 检查用户名与密码
	if len(user.Nickname) == 0 || len(user.Password) == 0 {
		return "", errors.New("用户名或密码不能为空！")
	}
	// 检查用户名是否重复
	if _, err = models.GetUserByName(user.Nickname); err == nil {
		return "", errors.New("用户名存在！")
	}
	// 准备数据
	u := models.User{
		Nickname: user.Nickname,
		Password: encodePassword(user.Password),
	}
	// 创建用户
	if err = models.CreateUser(&u); err != nil {
		log.Println(err.Error())
		return "", errors.New("注册失败！")
	}
	// 生成token
	if token, err = jwt.GenerateToken(int64(u.UID), u.Nickname); err != nil {
		log.Println(err.Error())
		return "", errors.New("token生成失败！")
	}
	return token, nil
}

func (service *UserServiceImpl) Logout(token string) (err error) {
	// 解析token
	claim, err := jwt.ParseToken(token)
	if err != nil {
		log.Println(err.Error())
		return errors.New("无效token！")
	}
	// 删除用户
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	if err = models.DeleteUserByID(uid); err != nil {
		log.Println(err.Error())
		return errors.New("注销失败！")
	}
	return nil
}

func (service *UserServiceImpl) Status(oldToken string) (newToken string, err error) {
	// 解析token
	claim, err := jwt.ParseToken(oldToken)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("无效token！")
	}
	// 生成新的token
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	if newToken, err = jwt.GenerateToken(uid, claim.Audience); err != nil {
		log.Println(err.Error())
		return "", errors.New("token生成失败！")
	}
	return newToken, nil
}

func (service *UserServiceImpl) GetUserList(page, size int) ([]*User, int64, error) {
	// 获取用户列表
	users, err := models.GetUserList(page, size)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取用户列表失败！")
	}
	var list []*User
	for _, u := range users {
		list = append(list, &User{
			ID:       u.UID,
			Nickname: u.Nickname,
		})
	}
	total, err := models.CountUser()
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取用户列表失败！")
	}
	return list, total, nil
}

func (service *UserServiceImpl) GetUser(uid int64) (*User, error) {
	// 获取用户
	u, err := models.GetUserById(uid)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取用户失败！")
	}
	return &User{
		ID:       u.UID,
		Nickname: u.Nickname,
	}, nil
}

func encodePassword(password string) string {
	return utile.EnCoder(password)
}
