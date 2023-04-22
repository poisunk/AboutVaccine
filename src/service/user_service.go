package service

import (
	"about-vaccine/src/entity"
	"about-vaccine/src/middleware/jwt"
	"about-vaccine/src/repo"
	"about-vaccine/src/schama"
	"about-vaccine/src/utile"
	"errors"
	"log"
	"strconv"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) Login(username, password string) (string, error) {
	if len(username) == 0 || len(password) == 0 {
		return "", errors.New("用户名或密码为空")
	}
	u, _, err := s.userRepo.GetByName(username)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("用户不存在")
	}
	if s.encodePassword(password) != u.Password {
		return "", errors.New("密码错误")
	}
	token, err := jwt.GenerateToken(u.Uid, u.Nickname)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("token生成失败")
	}
	return token, nil
}

func (s *UserService) Register(username, password string) (string, error) {
	if len(username) == 0 || len(password) == 0 {
		return "", errors.New("用户名或密码为空")
	}
	u, _, err := s.userRepo.GetByName(username)
	if err == nil {
		log.Println(err.Error())
		return "", errors.New("用户已存在")
	}
	u = &entity.User{
		Nickname: username,
		Password: s.encodePassword(password),
	}
	err = s.userRepo.Create(u)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("注册失败")
	}
	token, err := jwt.GenerateToken(u.Uid, u.Nickname)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("token生成失败")
	}
	return token, nil
}

func (s *UserService) Logout(token string) error {
	claim, err := jwt.ParseToken(token)
	if err != nil {
		log.Println(err.Error())
		return errors.New("无效的token")
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	err = s.userRepo.Delete(uid)
	if err != nil {
		log.Println(err.Error())
		return errors.New("注销失败")
	}
	return nil
}

func (s *UserService) LoginWithToken(token string) (string, error) {
	claim, err := jwt.ParseToken(token)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("无效的token")
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	newToken, err := jwt.GenerateToken(uid, claim.Audience)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("token生成失败")
	}
	return newToken, nil
}

func (s *UserService) SearchUserByName(keyword string, page, pageSize int) ([]*schama.User, int64, error) {
	u, total, err := s.userRepo.GetBySimilarName(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	list := make([]*schama.User, len(u))
	for i, v := range u {
		list[i] = &schama.User{}
		list[i].GetFormEntity(v)
	}
	return list, total, nil
}

func (s *UserService) encodePassword(password string) string {
	return utile.EnCoder(password)
}
