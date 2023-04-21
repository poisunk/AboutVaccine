package entity

type User struct {
	Uid      int    `json:"uid"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}
