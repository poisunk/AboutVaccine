package schama

import "about-vaccine/src/entity"

type User struct {
	Uid      int64  `json:"uid"`
	Nickname string `json:"nickname"`
}

type UserClaim struct {
	Uid      int64  `json:"uid"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

func (u *User) GetFormEntity(user *entity.User) {
	u.Uid = user.Uid
	u.Nickname = user.Nickname
}
