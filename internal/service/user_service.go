package service

import (
	"about-vaccine/internal/middleware/jwt"
	"about-vaccine/internal/schema"
	"about-vaccine/internal/service/user"
	"about-vaccine/internal/utile"
	"errors"
	"log"
	"strconv"
)

type UserService struct {
	common *user.UserCommon
}

func NewUserService(userCommon *user.UserCommon) *UserService {
	return &UserService{
		common: userCommon,
	}
}

func (s *UserService) Login(username, password string) (*schema.UserClaim, error) {
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("用户名或密码为空")
	}
	u, _, err := s.common.GetByName(username)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("用户不存在")
	}
	if s.encodePassword(password) != u.Password {
		return nil, errors.New("密码错误")
	}
	token, err := jwt.GenerateToken(u.Id, u.Nickname)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("token生成失败")
	}
	claim := &schema.UserClaim{
		UserInfo: schema.UserInfo{
			Uid:      u.Id,
			Nickname: u.Nickname,
		},
		Token: token,
	}
	return claim, nil
}

func (s *UserService) Register(username, password string) (*schema.UserClaim, error) {
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("用户名或密码为空")
	}
	has, _ := s.common.CheckName(username)
	if has {
		return nil, errors.New("用户已存在")
	}
	u := &schema.UserAdd{
		Nickname: username,
		Password: s.encodePassword(password),
	}
	uid, err := s.common.Create(u)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("注册失败")
	}
	token, err := jwt.GenerateToken(uid, u.Nickname)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("token生成失败")
	}
	claim := &schema.UserClaim{
		UserInfo: schema.UserInfo{
			Uid:      uid,
			Nickname: u.Nickname,
		},
		Token: token,
	}
	return claim, nil
}

func (s *UserService) Logout(token string) error {
	claim, err := jwt.ParseToken(token)
	if err != nil {
		log.Println(err.Error())
		return errors.New("无效的token")
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	err = s.common.Delete(uid)
	if err != nil {
		log.Println(err.Error())
		return errors.New("注销失败")
	}
	return nil
}

func (s *UserService) LoginWithToken(token string) (*schema.UserClaim, error) {
	claim, err := jwt.ParseToken(token)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("无效的token")
	}
	has, _ := s.common.CheckName(claim.Audience)
	if !has {
		return nil, errors.New("用户不存在")
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	newToken, err := jwt.GenerateToken(uid, claim.Audience)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("token生成失败")
	}
	userClaim := &schema.UserClaim{
		UserInfo: schema.UserInfo{
			Uid:      uid,
			Nickname: claim.Audience,
		},
		Token: newToken,
	}
	return userClaim, nil
}

func (s *UserService) SearchUserByName(keyword string, page, pageSize int) ([]*schema.UserInfo, int64, error) {
	list, total, err := s.common.GetListBySimilarName(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	return list, total, nil
}

func (s *UserService) encodePassword(password string) string {
	return utile.EnCoder(password)
}
