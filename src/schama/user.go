package schama

import "about-vaccine/src/entity"

type User struct {
	Uid      int64
	Nickname string
}

func (u *User) GetFormEntity(user *entity.User) {
	u.Uid = user.Uid
	u.Nickname = user.Nickname
}
