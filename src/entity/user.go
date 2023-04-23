package entity

type User struct {
	Uid      int64  `xorm:"autoincr" json:"uid"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}
